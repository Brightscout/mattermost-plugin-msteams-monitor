import {createSlice, PayloadAction} from '@reduxjs/toolkit';

const initialState: WebsocketEventState = {
    isConnected: false,
};

export const websocketEventSlice = createSlice({
    name: 'websocketEventSlice',
    initialState,
    reducers: {
        toggleIsConnected: (state: WebsocketEventState, action: PayloadAction<boolean>) => {
            state.isConnected = action.payload;
        },
    },
});

export const {toggleIsConnected} = websocketEventSlice.actions;

export default websocketEventSlice.reducer;
