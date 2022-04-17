import React from 'react';
import './App.css';
import { Route, Switch } from 'wouter';
import { AllRoutes } from './routes/routes';

import Navbar from './components/Navbar';

function App(): JSX.Element {
  return (
    <div className="App">
      <Navbar />
      <div className="app-content min-h-screen bg-gray-900 text-white">
      <Switch>
          { AllRoutes.map(({ label, path, component: Component}, props) => (
            <Route
              key={label}
              path={path}
              component={Component}
              {...props}
            />
          ))}
        </Switch>
      </div>
    </div>
  );
}

export default App;
