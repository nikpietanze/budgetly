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
        <nav class={`navbar bg-base-100 border-b ${theme() === Themes.dark ? "border-slate-700" : "border-slate-300"} px-6`}>
            <div class="navbar-start">
                <div class="dropdown">
                    <label tabindex="0" class="btn btn-ghost btn-circle">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h7" /></svg>
                    </label>
                    <ul tabindex="0" class="menu menu-normal dropdown-content mt-3 p-2 shadow-md bg-base-100 rounded-lg w-52">
                        <li><a href="/">Home</a></li>
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
                <Show
                    when={!loading() && user()}
                    fallback={<button class="btn btn-ghost" onClick={() => loginWithRedirect()}>Sign in</button>}
                >
                    <div class="dropdown dropdown-end">
                        <label tabindex="0" class="btn btn-ghost btn-circle flex justify-center items-center">
                            <img class="rounded-full shadow shadow-stone-900 h-10 w-10" src={user().picture} />
                        </label>
                        <ul tabindex="0" class="menu menu-normal dropdown-content mt-3 p-2 shadow-md bg-base-100 rounded-lg w-52">
                            <li><span onClick={() => logout()}>Sign out</span></li>
                        </ul>
                    </div>
                </Show>
            </div>
        </nav >
    );
};

export default Navbar;
