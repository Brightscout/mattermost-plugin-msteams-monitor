import {
    CHANNEL_HEADER_BUTTON_TOOLTIP,
    RIGHT_SIDEBAR_HEADER,
    HEADER_CSRF_TOKEN,
    MMCSRF,
} from './common';

import {error} from './messages';

import {pluginApiServiceConfigs} from './apiService';

export default {
    common: {
        RIGHT_SIDEBAR_HEADER,
        CHANNEL_HEADER_BUTTON_TOOLTIP,
        HEADER_CSRF_TOKEN,
        MMCSRF,
    },
    messages: {
        error,
    },
    pluginApiServiceConfigs,
};
