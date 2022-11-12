import { Component, createEffect } from 'solid-js';
import { useAuth0 } from '../../Auth0';

const Logout: Component = () => {
    const { isLoading, logout } = useAuth0();

    createEffect(() => {
        if (!isLoading()) {
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
