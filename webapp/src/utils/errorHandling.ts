/**
 * Custom error handling
 */
import pluginConstants from 'pluginConstants';

const getMeCustomError = (errorState?: ApiErrorResponse) => {
    return null;
};

const customErrorAndModuleMap: Partial<Record<CustomErrorModules, (errorState?: ApiErrorResponse) => string | null>> = {
    getMe: (errorState?: ApiErrorResponse) => getMeCustomError(errorState),
};

const getErrorMessage = (
    isError: boolean,
    module: CustomErrorModules,
    errorState?: ApiErrorResponse,
): string => {
    if (!isError) {
        return '';
    }

    // Return custom or API error for modules
    return customErrorAndModuleMap[module]?.(errorState) ?? errorState?.data.error ?? pluginConstants.messages.error.GENERIC;
};

export default getErrorMessage;
