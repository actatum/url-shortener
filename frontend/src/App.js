import React from 'react';
import './App.css';
import { Switch, Route } from "react-router-dom"
import Home from "./pages/home.component"
import Slug from "./components/slug/slug.component"

function App() {
  return (
    <div className="App">
      <Switch>
        <Route exact path="/" component={Home} />
        <Route path="/:id" component={Slug} />
      </Switch>
    </div>
  );
}

export default App;
