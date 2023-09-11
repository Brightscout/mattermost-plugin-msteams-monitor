import {pluginId} from 'manifest';

import getErrorMessage from './errorHandling';

// TODO: Use Mattermost's site URL from the redux store
const getBaseUrls = (mmSiteUrl: string): {pluginApiBaseUrl: string; mattermostApiBaseUrl: string} => {
    const pluginUrl = `${mmSiteUrl}/plugins/${pluginId}`;
    const pluginApiBaseUrl = `${pluginUrl}/api/v1`;
    const mattermostApiBaseUrl = `${mmSiteUrl}/api/v4`;

    return {pluginApiBaseUrl, mattermostApiBaseUrl};
};

export const getCommandArgs = (command: string) => {
    const myRegexp = /[^\s"]+|"([^"]*)"/gi;
    const myArray = [];
    let match;
    do {
        match = myRegexp.exec(command);
        if (match != null) {
            myArray.push(match[1] ?? match[0]);
        }
    } while (match != null);
    return myArray.length > 2 ? myArray.slice(2) : [];
};

export default {
    getBaseUrls,
    getErrorMessage,
};
