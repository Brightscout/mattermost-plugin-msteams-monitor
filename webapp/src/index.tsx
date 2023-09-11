import React from 'react';
import Cookies from 'js-cookie';

import {Store, Action} from 'redux';

import {GlobalState} from 'mattermost-redux/types/store';

// eslint-disable-next-line import/no-unresolved
import {PluginRegistry} from 'types/mattermostWebapp';

// Constants
import pluginConstants from 'pluginConstants';

// Containers
import ChannelHeaderButton from 'containers/components/channelHeaderButton';
import RHS from 'containers/rhs';

// Reducer
import reducer from 'reducers';

import {pluginId} from './manifest';

export default class Plugin {
    // eslint-disable-next-line @typescript-eslint/no-unused-vars, @typescript-eslint/no-empty-function
    public async initialize(registry: PluginRegistry, store: Store<GlobalState, Action<Record<string, unknown>>>) {
        // @see https://developers.mattermost.com/extend/plugins/webapp/reference/
        registry.registerReducer(reducer);

        const {showRHSPlugin} = registry.registerRightHandSidebarComponent(RHS, pluginConstants.common.RIGHT_SIDEBAR_HEADER);
        registry.registerChannelHeaderButtonAction(<ChannelHeaderButton/>, () => store.dispatch(showRHSPlugin), null, pluginConstants.common.CHANNEL_HEADER_BUTTON_TOOLTIP);
    }
}

declare global {
    interface Window {
        registerPlugin(id: string, plugin: Plugin): void
    }
}

window.registerPlugin(pluginId, new Plugin());
