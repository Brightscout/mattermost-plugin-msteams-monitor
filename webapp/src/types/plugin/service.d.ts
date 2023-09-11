type PluginApiServiceName = 'getMe';

type ApiErrorResponse = {
    data: {
        error: string,
    },
    status: number,
}

type CustomErrorModules = 'getMe';

type PluginApiService = {
    path: string,
    method: HttpMethod,
    apiServiceName: PluginApiServiceName,
}

type APIRequestPayload = void;
