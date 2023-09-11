// eslint-disable-next-line import/no-unresolved
import {BaseQueryApi} from '@reduxjs/toolkit/dist/query/baseQueryTypes';
import {createApi, fetchBaseQuery} from '@reduxjs/toolkit/query/react';
import Cookies from 'js-cookie';
import {GlobalState} from 'mattermost-redux/types/store';

import {UserProfile} from 'mattermost-redux/types/users';

import Constants from 'pluginConstants';

import Utils from 'utils';

const handleBaseQuery = async (
    args: {
        url: string,
        method: string,
    },
    api: BaseQueryApi,
    extraOptions: Record<string, string> = {},
) => {
    const globalReduxState = api.getState() as GlobalState;
    const result = await fetchBaseQuery({
        baseUrl: Utils.getBaseUrls(globalReduxState?.entities?.general?.config?.SiteURL as string).pluginApiBaseUrl,
        prepareHeaders: (headers) => {
            headers.set(Constants.common.HEADER_CSRF_TOKEN, Cookies.get(Constants.common.MMCSRF) ?? '');

            return headers;
        },
    })(
        args,
        api,
        extraOptions,
    );
    return result;
};

// Service to make plugin API requests
export const examplePluginApi = createApi({
    reducerPath: 'examplePluginApi',
    baseQuery: handleBaseQuery,
    endpoints: (builder) => ({
        [Constants.pluginApiServiceConfigs.getMe.apiServiceName]: builder.query<UserProfile, void>({
            query: () => ({
                url: Constants.pluginApiServiceConfigs.getMe.path,
                method: Constants.pluginApiServiceConfigs.getMe.method,
            }),
        }),
    }),
});
