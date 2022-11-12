import { createAuth0Client } from '@auth0/auth0-spa-js';
import {
    createContext,
    useContext,
    createSignal,
    createResource,
    JSX
} from 'solid-js';
import type { Auth0State } from './types/auth0';

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
    const [isAuthenticated, setIsAuthenticated] = createSignal<boolean>(false);
    const [isLoading, setIsLoading] = createSignal<boolean>(true);
    const [user, setUser] = createSignal(null);
    const [error, setError] = createSignal<Error>(null);

    if (!props.onRedirectCallback) {
        window.history.replaceState({}, document.title, window.location.pathname);
    }

    const auth0ClientPromise = createAuth0Client({
        domain: props.domain,
        clientId: props.clientId,
        authorizationParams: {
            redirect_uri: props.redirectUri,
        }
    });

    createResource(async () => {
        const client = await auth0ClientPromise;

        try {
            const search = window.location.search;
            if (
                (search.includes("code=") || search.includes("error=")) &&
                search.includes("state=")
            ) {
                const { appState } = await client.handleRedirectCallback();
                props.onRedirectCallback(appState);
            }
        } catch (err) {
            setError(err);
        } finally {
            setIsAuthenticated(await client.isAuthenticated());
            setUser(await client.getUser() || null);
            setIsLoading(false);
        }
    });

    const login = async () => {
        try {
            const client = await auth0ClientPromise;
            await client.loginWithRedirect();
        } catch (err) {
            setError(err);
        }
    }

    const logout = async () => {
        try {
            const client = await auth0ClientPromise;
            await client.logout({
                async onRedirect(url) {
                    window.location.replace(url);
                }
            });
        } catch (err) {
            setError(err);
        }
    }

    const getAccessToken = async () => {
        try {
            const client = await auth0ClientPromise;
            return await client.getTokenSilently();
        } catch (err) {
            setError(err);
        }
    }

    return (
        <Auth0Context.Provider
            value={{
                isAuthenticated,
                isLoading,
                user,
                error,
                login,
                logout,
                getAccessToken
            }}
        >
            {props.children}
        </Auth0Context.Provider>
    );
}

function isRedirect(url: string) {
    const [, query] = url.split('?');
    return query && query.includes('code=') && query.includes('state=');
}

