package plugin

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/mattermost/mattermost-plugin-api/experimental/command"
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/plugin"

	"github.com/brightscout/mattermost-plugin-msteams-monitor/server/constants"
)

type HandlerFunc func(p *Plugin, c *plugin.Context, commandArgs *model.CommandArgs, args ...string) (*model.CommandResponse, *model.AppError)

type Handler struct {
	handlers       map[string]HandlerFunc
	defaultHandler HandlerFunc
}

var exampleCommandHandler = Handler{
	handlers: map[string]HandlerFunc{
		constants.CommandHelp:    helpCommand,
		constants.CommandConnect: connectCommand,
	},
	defaultHandler: executeDefault,
}

// Handle function calls the respective handlers of the commands.
// It checks whether any HandlerFunc is present for the given command by checking in the "exampleCommandHandler".
// If the command is present, it calls its handler function, else calls the default handler.
func (ch *Handler) Handle(p *Plugin, c *plugin.Context, commandArgs *model.CommandArgs, args ...string) (*model.CommandResponse, *model.AppError) {
	for arg := len(args); arg > 0; arg-- {
		handler := ch.handlers[strings.Join(args[:arg], "/")]
		if handler != nil {
			return handler(p, c, commandArgs, args[arg:]...)
		}
	}
	return ch.defaultHandler(p, c, commandArgs, args...)
}

func (p *Plugin) getAutoCompleteData() *model.AutocompleteData {
	cmd := model.NewAutocompleteData(constants.CommandTriggerName, "[command]", fmt.Sprintf("Available commands: %s", constants.CommandHelp))

	connect := model.NewAutocompleteData("connect", "", "Connect your Mattermost account to your MS Teams account")
	cmd.AddCommand(connect)

	help := model.NewAutocompleteData(constants.CommandHelp, "", fmt.Sprintf("Show %s slash command help", constants.CommandTriggerName))
	cmd.AddCommand(help)

	return cmd
}

func (p *Plugin) getCommand() (*model.Command, error) {
	iconData, err := command.GetIconData(p.API, "public/assets/example-bot.svg")
	if err != nil {
		return nil, errors.Wrap(err, "failed to get example icon")
	}

	return &model.Command{
		Trigger:              constants.CommandTriggerName,
		AutoComplete:         true,
		AutoCompleteDesc:     fmt.Sprintf("Available commands: %s", constants.CommandHelp),
		AutoCompleteHint:     "[command]",
		AutocompleteData:     p.getAutoCompleteData(),
		AutocompleteIconData: iconData,
	}, nil
}

func helpCommand(p *Plugin, _ *plugin.Context, commandArgs *model.CommandArgs, _ ...string) (*model.CommandResponse, *model.AppError) {
	p.sendEphemeralPost(commandArgs, constants.HelpText)
	return &model.CommandResponse{}, nil
}

// TODO: Complete connect command during oAuth completion
func connectCommand(p *Plugin, _ *plugin.Context, commandArgs *model.CommandArgs, _ ...string) (*model.CommandResponse, *model.AppError) {
	connectURL := "connect URL"
	p.sendEphemeralPost(commandArgs, fmt.Sprintf("[Click here to connect your account](%s)", connectURL))
	return &model.CommandResponse{}, nil
}

func executeDefault(p *Plugin, _ *plugin.Context, commandArgs *model.CommandArgs, _ ...string) (*model.CommandResponse, *model.AppError) {
	p.sendEphemeralPost(commandArgs, constants.InvalidCommand+constants.HelpText)
	return &model.CommandResponse{}, nil
}

// Handles executing a slash command
func (p *Plugin) ExecuteCommand(c *plugin.Context, commandArgs *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	args := strings.Fields(commandArgs.Command)

	return exampleCommandHandler.Handle(p, c, commandArgs, args[1:]...)
}
