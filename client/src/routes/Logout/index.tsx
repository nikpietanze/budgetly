import { Component, onMount } from 'solid-js';
import { useAuth0 } from '../../Auth0';

const Logout: Component = () => {
    onMount(() => {
        logout();
    });

    return (
        <>
            <h1>Logging out...</h1>
        </>
    )
}

export default Logout;
