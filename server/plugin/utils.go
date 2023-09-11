package plugin

import (
	"errors"

	"github.com/mattermost/mattermost-server/v6/model"
)

var ErrNotFound = errors.New("not found")

// sendEphemeralPost sends an ephermal message
func (p *Plugin) sendEphemeralPost(args *model.CommandArgs, text string) {
	post := &model.Post{
		UserId:    p.botUserID,
		ChannelId: args.ChannelId,
		Message:   text,
	}
	_ = p.API.SendEphemeralPost(args.UserId, post)
}
