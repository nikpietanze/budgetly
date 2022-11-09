import type { Component } from 'solid-js';
import { useAuth0 } from '../../Auth0';

const Home: Component = () => {
    const { isAuthenticated } = useAuth0();
    const auth = isAuthenticated();
    console.log('authenticated', isAuthenticated())
    return (
        <>
            <div class="hero py-44 bg-base-200">
                <div class="hero-content text-center">
                    <div class="max-w-xl">
                        <h1 class="text-5xl font-bold">Welcome to Budgetly</h1>
                        <p class="py-6">Provident cupiditate voluptatem et in. Quaerat fugiat ut assumenda excepturi exercitationem quasi. In deleniti eaque aut repudiandae et a id nisi.</p>
                        <a class="btn btn-primary" href="/login">Get Started</a>
                    </div>
                </div>
            </div>
            <div>
                <h2>Authenticated: {auth}</h2>
            </div>
        </>
    )
}

export default Home;
