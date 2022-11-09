import { Component, createEffect, onMount } from 'solid-js';
import { useAuth0 } from '../../Auth0';

const Logout: Component = () => {
    const { loading, logout } = useAuth0();

    createEffect(() => {
        if (!loading()) {
            logout();
        }
    });

    return (
        <>
            <h1>Logging out...</h1>
        </>
    )
}

export default Logout;
