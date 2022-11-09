import type { Component } from 'solid-js';
import { createSignal, createEffect, Show } from 'solid-js';
import logoLight from '../assets/img/budgetly_light.png';
import logoDark from '../assets/img/budgetly_dark.png';
import { useAuth0 } from '../Auth0';
import type { User } from '@auth0/auth0-spa-js';

enum Themes {
    dark = "dark",
    light = "light",
}

const Navbar: Component = () => {
    const { loading, isAuthenticated, loginWithRedirect, logout } = useAuth0();
    const [user, setUser] = createSignal<User>();
    const [theme, setTheme] = createSignal(Themes.dark);

    createEffect(async () => {
        if (!loading()) {
            if (isAuthenticated()) {
                setUser(useAuth0().user());
            };
        }
    })

    const savedTheme = localStorage.getItem('theme');
    setTheme(savedTheme === Themes.dark ? Themes.dark : Themes.light);

    function toggleTheme(theme: Themes) {
        const htmlEl = document.querySelector('html');
        const logoEl: HTMLImageElement | null = document.querySelector("nav .navbar-center img");

        if (htmlEl && logoEl) {
            if (theme === Themes.light) {
                htmlEl.classList.remove(Themes.dark);
                htmlEl.classList.add(Themes.light);
                htmlEl.setAttribute("data-theme", Themes.light);
                logoEl.src = logoDark;
                localStorage.setItem('theme', Themes.light);
            } else {
                htmlEl.classList.add(Themes.dark);
                htmlEl.classList.remove(Themes.light);
                htmlEl.setAttribute("data-theme", Themes.dark);
                logoEl.src = logoLight;
                localStorage.setItem('theme', Themes.dark);
            }
        }
        return;
    }

    createEffect(() => toggleTheme(theme()));

    return (
        <nav class="navbar bg-base-100 border-b border-slate-700">
            <div class="navbar-start">
                <div class="dropdown flex">
                    <label tabindex="0" class="btn btn-ghost btn-circle ml-2">
                        <Show
                            when={isAuthenticated()}
                            fallback={
                                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h7" /></svg>
                            }
                        >
                            <img class="rounded-full" src={user()?.picture} />
                        </Show>
                    </label>
                    <ul tabindex="0" class="menu menu-compact dropdown-content mt-14 p-2 shadow-lg bg-base-100 rounded-lg w-52">
                        <li><a href="/">Home</a></li>
                        <Show
                            when={!loading() && !isAuthenticated()}
                            fallback={
                                <li><span onClick={() => logout()}>Logout</span></li>
                            }
                        >
                            <li><span onClick={() => loginWithRedirect()}>Login</span></li>
                        </Show>
                    </ul>
                </div>
            </div>
            <div class="navbar-center">
                <a class="btn btn-ghost normal-case text-xl h-auto py-2 px-4" href="/">
                    <img src={logoLight} class="w-44" alt="Budgetly" />
                </a>
            </div>
            <div class="navbar-end">
                <div class="theme-selector flex gap-2 mr-2">
                    {/* @ts-ignore */}
                    <input type="checkbox" class="toggle toggle-md bg-slate-600" checked={savedTheme && savedTheme === Themes.light} onClick={(e) => setTheme(e.target.checked ? Themes.light : Themes.dark)} />
                    <svg class="h-6" viewBox="0 0 24 24" style="fill: none; stroke: currentcolor;"><path stroke-linecap="round" stroke-linejoin="round" d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z"></path></svg>
                </div>
                <button class="btn btn-ghost btn-circle">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" /></svg>
                </button>
                <button class="btn btn-ghost btn-circle">
                    <div class="indicator">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" /></svg>
                        <span class="badge badge-xs badge-primary indicator-item"></span>
                    </div>
                </button>
            </div>
        </nav >
    );
};

export default Navbar;
