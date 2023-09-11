package main

import (
	mattermostPlugin "github.com/mattermost/mattermost-server/v6/plugin"

	"github.com/brightscout/mattermost-plugin-msteams-monitor/server/plugin"
)

func main() {
	mattermostPlugin.ClientMain(&plugin.Plugin{})
}
