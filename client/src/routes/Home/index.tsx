import type { Component } from 'solid-js';

const Home: Component = () => {
    return (
        <div class="hero py-44 bg-base-200">
            <div class="hero-content text-center">
                <div class="max-w-xl">
                    <h1 class="text-5xl font-bold">Welcome to Budgetly</h1>
                    <p class="py-6">Provident cupiditate voluptatem et in. Quaerat fugiat ut assumenda excepturi exercitationem quasi. In deleniti eaque aut repudiandae et a id nisi.</p>
                    <a class="btn btn-primary" href="/login">Get Started</a>
                </div>
            </div>
        </div>
    )
}

export default Home;
