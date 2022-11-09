import {
    PopupLoginOptions,
    RedirectLoginOptions,
    GetTokenSilentlyOptions,
    GetTokenWithPopupOptions,
    LogoutOptions,
    IdToken
} from '@auth0/auth0-spa-js';

export interface Auth0State {
    isAuthenticated: () => boolean;
    user: () => any;
    loading: () => boolean;
    popupOpen: () => boolean;
    loginWithPopup: (options: PopupLoginOptions) => Promise<void>;
    handleRedirectCallback: () => Promise<void>;
    getIdTokenClaims: () => Promise<IdToken>;
    loginWithRedirect: (options?: RedirectLoginOptions) => Promise<void>;
    getTokenSilently: (options?: GetTokenSilentlyOptions) => Promise<string>;
    getTokenWithPopup: (options?: GetTokenWithPopupOptions) => Promise<string>;
    logout: (options?: LogoutOptions) => Promise<void>;
}
