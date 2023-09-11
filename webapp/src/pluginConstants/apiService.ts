// Plugin api service (RTK query) configs
export const pluginApiServiceConfigs: Record<PluginApiServiceName, PluginApiService> = {
    getMe: {
        path: '/user/me',
        method: 'GET',
        apiServiceName: 'getMe',
    },
};
