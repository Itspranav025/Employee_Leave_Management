import React from 'react';
import './App.css';
import FrontendTask from './frontendtask';

function App() {
  return (
    <div className="app-container">
      <header className="app-header">
        <h1>Leave Form App</h1>
      </header>
      <main className="app-main">
        <section className="app-section">
          <FrontendTask />
        </section>
      </main>
      <footer className="app-footer">
        <p>&copy; {new Date().getFullYear()} Your Company. All rights reserved.</p>
      </footer>
    </div>
  );
}

export default App;
