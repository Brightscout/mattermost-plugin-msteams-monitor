export const getRhsState = (state: ReduxState): {isSidebarOpen: boolean} => state.views.rhs;

// TODO: Configure and use plugin id as a constant
// Plugin state
const getPluginState = (state: ReduxState) => state['plugins-mattermost-plugin-template'];

export const getApiRequestCompletionState = (state: ReduxState): ApiRequestCompletionState => getPluginState(state).apiRequestCompletionSlice;
