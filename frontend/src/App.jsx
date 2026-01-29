/*
App Component - Loading screen before redirecting to WhatsApp Web
*/
import { useEffect } from 'react';
import './App.css';

function App() {
    useEffect(() => {
        const timer = setTimeout(() => {
            window.location.href = "https://web.whatsapp.com";
        }, 800);
        return () => clearTimeout(timer);
    }, []);

    return (
        <div id="App" className="loading-container">
            <div className="loader-dots">
                <span></span>
                <span></span>
                <span></span>
            </div>
            <h2 className="loading-text">Loading WhatsApp</h2>
        </div>
    );
}

export default App;
