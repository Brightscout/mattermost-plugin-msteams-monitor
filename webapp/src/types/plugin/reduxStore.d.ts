interface ReduxState extends GlobalState {
    views: {
        rhs: {
            isSidebarOpen: boolean
        }
    }
    'plugins-mattermost-plugin-template': RootState<{ [x: string]: QueryDefinition<void, BaseQueryFn<string | FetchArgs, unknown, FetchBaseQueryError, {}, FetchBaseQueryMeta>, never, void, 'examplePluginApi'>; }, never, 'examplePluginApi'>
}

type ApiRequestCompletionState = {
    requests: PluginApiServiceName[];
}

type GlobalModalState = {
    modalId: ModalId;
    commandArgs: Array<string>;
}

type WebsocketEventState = {
    isConnected: boolean;
};
