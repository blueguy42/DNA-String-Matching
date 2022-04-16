import React from 'react';
import logo from './logo.svg';
import './App.css';

import Navbar from './components/Navbar';
import Insert from "./pages/Insert";

function App(): JSX.Element {
  return (
    <div className="App">
      <Navbar />
      {/* <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header> */}

      < Insert />
    </div>
  );
}

export default App;
