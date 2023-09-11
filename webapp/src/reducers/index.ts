import {combineReducers} from 'redux';

import {examplePluginApi} from 'services';

import globalModalSlice from './globalModal';
import apiRequestCompletionSlice from './apiRequest';
import websocketEventSlice from './websocketEvent';

const reducers = combineReducers({
    apiRequestCompletionSlice,
    globalModalSlice,
    websocketEventSlice,
    [examplePluginApi.reducerPath]: examplePluginApi.reducer,
});

export default reducers;
