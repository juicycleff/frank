import {authRefreshToken, client as fclient, type Client} from '../../sdk/index';
import {getConfig} from '../config';
import {clearTokenData, getTokenData, isTokenExpired, setTokenData} from './token';
import {TokenData} from '../types';

// type Client = () => typeof fclient;

// Create a base API client
export const createApiClient: Client = () => {
    const config = getConfig();

    fclient.setConfig({
        baseUrl: config.baseUrl,
        headers: {}
    });

    return fclient;
};

// Create an authenticated API client with token handling
export const createAuthenticatedClient: Client = () => {
    const config = getConfig();
    const tokenData = getTokenData();

    fclient.setConfig({
        baseUrl: config.baseUrl,
        headers: tokenData ? {
            'Authorization': `Bearer ${tokenData.token}`
        } : {}
    });

    // Add organization ID to headers if provided
    if (config.organizationId) {
        fclient.setConfig({
            ...fclient.getConfig(),
            headers: {
                ...fclient.getConfig().headers,
                'X-Organization-ID': config.organizationId
            }
        });
    }

    return fclient;
};

// Function to refresh the token
export const refreshAuthToken = async (): Promise<TokenData | null> => {
    const tokenData = getTokenData();
    if (!tokenData) return null;

    try {
        const client = createApiClient();
        const { data } = await authRefreshToken({
            body: {
                refresh_token: tokenData.refreshToken
            },
            throwOnError: true,
            client,
        });

        const newTokenData: TokenData = {
            token: data.token,
            refreshToken: data.refresh_token,
            expiresAt: Number(data.expires_at),
        };

        setTokenData(newTokenData);
        return newTokenData;
    } catch (error) {
        clearTokenData();
        return null;
    }
};

// Get authorized client with automatic token refresh
export const getAuthClient = async (): Promise<Client> => {
    if (isTokenExpired() && getTokenData()) {
        await refreshAuthToken();
    }

    return createAuthenticatedClient() as any;
};