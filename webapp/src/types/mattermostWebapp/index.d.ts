export interface PluginRegistry {
    registerSlashCommandWillBePostedHook(hook: (message: string, args: MmHookArgTypes) => Promise<({message?: string, args?: MmHookArgTypes})>)
    registerReducer(reducer);
    registerChannelHeaderButtonAction(icon: JSX.Element, action: () => void, dropdownText: string | null, tooltipText: string | null);
    registerRightHandSidebarComponent(component: () => JSX.Element, title: string | JSX.Element);
    registerWebSocketEventHandler(event: string, handler: (msg: WebsocketEventParams) => void)

    // Add more if needed from https://developers.mattermost.com/extend/plugins/webapp/reference
}

type MmHookArgTypes = {
    channel_id: string,
    team_id: string,
    root_id: string
}
