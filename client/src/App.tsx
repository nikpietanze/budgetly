import type { Component } from 'solid-js';
import { Routes, Route } from '@solidjs/router';
import Navbar from './components/Navbar';
import Footer from './components/Footer';
import Home from './routes/Home';
import Login from './routes/Login';
import Account from './routes/Account';
import { useAuth0 } from './Auth0';

const App: Component = () => {
    const { isAuthenticated } = useAuth0();
    return (
        <div class="flex flex-col min-h-screen">
            <Navbar />
            <main class="flex-auto">
                <Routes>
                    <Route path="/" component={Home} />
                    <Route path="/login" component={Login} />
                    <Route path="/account" component={Account} />
                </Routes>
            </main>
            <Footer />
        </div>
    );
};

export default App;
