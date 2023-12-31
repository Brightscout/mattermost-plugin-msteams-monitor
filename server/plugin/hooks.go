package plugin

import (
	"github.com/pkg/errors"

	pluginapi "github.com/mattermost/mattermost-plugin-api"

	"github.com/mattermost/mattermost-plugin-template/server/config"
)

// Invoked when configuration changes may have been made.
func (p *Plugin) OnConfigurationChange() error {
	var configuration = new(config.Configuration)

	// Load the public configuration fields from the Mattermost server configuration.
	if err := p.API.LoadPluginConfiguration(configuration); err != nil {
		return errors.Wrap(err, "failed to load plugin configuration")
	}

	if err := configuration.ProcessConfiguration(); err != nil {
		p.API.LogError("Error in processing configuration.", "Error", err.Error())
		return err
	}

	if err := configuration.IsValid(); err != nil {
		p.API.LogError("Error in validating configuration.", "Error", err.Error())
		return err
	}

	p.setConfiguration(configuration)

	return nil
}

// Invoked when the plugin is activated
func (p *Plugin) OnActivate() error {
	p.client = pluginapi.NewClient(p.API, p.Driver)

	if err := p.OnConfigurationChange(); err != nil {
		return err
	}

	if err := p.initBotUser(); err != nil {
		return err
	}

	command, err := p.getCommand()
	if err != nil {
		return errors.Wrap(err, "failed to get command")
	}

	err = p.API.RegisterCommand(command)
	if err != nil {
		return errors.Wrap(err, "failed to register command")
	}

	p.router = p.InitAPI()
	p.InitRoutes()
	return nil
}
