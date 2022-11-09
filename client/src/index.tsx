import { render } from 'solid-js/web';
import { Auth0Provider } from './Auth0';
import { Router } from "@solidjs/router";
import history from './history';

import './index.css';
import App from './App';

const onRedirectCallback = appState => {
  history.push(
    appState && appState.targetUrl
      ? appState.targetUrl
      : window.location.pathname
  );
};

render(
    () => (
        <Auth0Provider
            domain={import.meta.env.VITE_AUTH0_DOMAIN}
            clientId={import.meta.env.VITE_AUTH0_CLIENT_ID}
            redirectUri={window.location.origin}
            onRedirectCallback={onRedirectCallback}
        >
            <Router>
                <App />
            </Router>
        </Auth0Provider>
    ),
    document.getElementById('root') as HTMLElement
);
