import React, {useEffect} from 'react';

import usePluginApi from 'hooks/usePluginApi';
import pluginConstants from 'pluginConstants';

const RHS = () => {
    const {getApiState, makeApiRequestWithCompletionStatus} = usePluginApi();

    useEffect(() => {
        makeApiRequestWithCompletionStatus(
            pluginConstants.pluginApiServiceConfigs.getMe.apiServiceName,
        );
    }, []);

    const {data, isLoading, isError, isSuccess} = getApiState(pluginConstants.pluginApiServiceConfigs.getMe.apiServiceName);

    const getApiData = () => {
        if (isLoading) {
            return <p>{'Loading...'}</p>;
        }
        if (isError) {
            return <p>{'Error<'}</p>;
        }
        if (isSuccess) {
            return <p>{'Email:'} {data?.email}</p>;
        }

        return '';
    };
    return (
        <>
            <h1>{'Hello!'}</h1>
            <p>{'This is a placeholder container for RHS'}</p>
            <h2>{'Below data is retreived from an API:'}</h2>
            {getApiData()}
        </>
    );
};

export default RHS;
