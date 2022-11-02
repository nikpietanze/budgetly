import type { Component } from 'solid-js';

const Footer: Component = () => {
    return (
        <footer class="footer footer-center p-4 bg-base-300 text-base-content">
            <div>
                <p>Copyright © {new Date().getFullYear()} - All right reserved by Budgetly</p>
            </div>
        </footer>
    )
}

export default Footer;
