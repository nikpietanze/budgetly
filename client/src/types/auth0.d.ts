import type {
    RedirectLoginOptions,
    GetTokenSilentlyOptions,
    LogoutOptions,
} from '@auth0/auth0-spa-js';

export interface Auth0State {
    isAuthenticated: () => boolean;
    isLoading: () => boolean;
    user: () => any;
    error: () => Error;
    login: (options?: RedirectLoginOptions) => Promise<void>;
    logout: (options?: LogoutOptions) => Promise<void>;
    getAccessToken: (options?: GetTokenSilentlyOptions) => Promise<string>;
}
