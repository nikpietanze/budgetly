import { createAuth0Client } from '@auth0/auth0-spa-js';
import type { Auth0Client } from '@auth0/auth0-spa-js';
import {
    onMount,
    createContext,
    useContext,
    createSignal,
    JSX
} from 'solid-js';
import type { Auth0State } from './types/auth0';

const DEFAULT_REDIRECT_CALLBACK = () => {
    window.history.replaceState({}, document.title, window.location.pathname);
}

export const Auth0Context = createContext<Auth0State>();
export const useAuth0 = () => useContext(Auth0Context);


export interface Auth0Props {
    children: JSX.Element;
    domain: string;
    clientId: string;
    redirectUri: string;
    onRedirectCallback?: (appState: any) => void;
}

export function Auth0Provider(props: Auth0Props) {
    const [isAuthenticated, setIsAuthenticated] = createSignal<
        boolean | undefined
    >(undefined);
    const [user, setUser] = createSignal();
    const [auth0Client, setAuth0Client] = createSignal<Auth0Client>();
    const [loading, setLoading] = createSignal<boolean>(true);
    const [popupOpen, setPopupOpen] = createSignal<boolean>(false);

    onMount(async () => {
        if (!props.onRedirectCallback) {
            props.onRedirectCallback = DEFAULT_REDIRECT_CALLBACK;
        }
        const auth0FromHook = await createAuth0Client({
            domain: props.domain,
            clientId: props.clientId,
        });
        setAuth0Client(auth0FromHook);

        if (window.location.search.includes("code=") &&
            window.location.search.includes("state=")) {
            const { appState } = await auth0FromHook.handleRedirectCallback();
            props.onRedirectCallback(appState);
        }

        const isAuthenticated = await auth0FromHook.isAuthenticated();
        setIsAuthenticated(isAuthenticated);
        if (isAuthenticated) {
            const user = await auth0FromHook.getUser();
            setUser(user);
        }

        setLoading(false);
    });

    const loginWithPopup = async (params = {}) => {
        setPopupOpen(true);

        try {
            await auth0Client().loginWithPopup(params);
        } catch (err) {
            console.error(err);
        } finally {
            setPopupOpen(false);
        }

        const user = await auth0Client().getUser();
        setUser(user);
        setIsAuthenticated(true);
    };

    const handleRedirectCallback = async () => {
        setLoading(true);
        await auth0Client().handleRedirectCallback();
        const user = await auth0Client().getUser();
        setLoading(false);
        setIsAuthenticated(true);
        setUser(user);
    };

    return (
        <Auth0Context.Provider
            value={{
                isAuthenticated,
                user,
                loading,
                popupOpen,
                loginWithPopup,
                handleRedirectCallback,
                getIdTokenClaims: () => auth0Client().getIdTokenClaims(),
                loginWithRedirect: (p) => auth0Client().loginWithRedirect(p),
                getTokenSilently: (p) => auth0Client().getTokenSilently(p),
                getTokenWithPopup: (p) => auth0Client().getTokenWithPopup(p),
                logout: (p) => auth0Client().logout(p)
            }}
        >
            {props.children}
        </Auth0Context.Provider>
    );
}

