import { createAuth0Client } from '@auth0/auth0-spa-js';
import type { Auth0Client, LogoutOptions } from '@auth0/auth0-spa-js';
import {
    createContext,
    useContext,
    createSignal,
    createResource,
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

    if (!props.onRedirectCallback) {
        props.onRedirectCallback = DEFAULT_REDIRECT_CALLBACK;
    }

    createResource(async () => {
        const client = await createAuth0Client({
            domain: props.domain,
            clientId: props.clientId,
        });
        setAuth0Client(client);
        console.log(auth0Client())

        if (window.location.search.includes("code=") &&
            window.location.search.includes("state=")) {
            const { appState } = await client.handleRedirectCallback();
            props.onRedirectCallback(appState);
        }

        const isAuthenticated = await client.isAuthenticated();
        setIsAuthenticated(isAuthenticated);
        if (isAuthenticated) {
            const user = await client.getUser();
            setUser(user);
        }

        setLoading(false);
        return client;
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
                loginWithRedirect: (opts) => auth0Client().loginWithRedirect(opts),
                getTokenSilently: (opts) => auth0Client().getTokenSilently(opts),
                getTokenWithPopup: (opts) => auth0Client().getTokenWithPopup(opts),
                logout: (opts) => auth0Client().logout(opts),
            }}
        >
            {props.children}
        </Auth0Context.Provider>
    );
}

