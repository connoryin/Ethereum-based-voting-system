import React from 'react';
import './App.css';
import Main from './pages/Main';
import Admin_homepage from './pages/Admin_homepage';
import Admin_login from './pages/Admin_login';
import IC_inputted from './pages/IC_inputted';
import Event_result from './pages/Event_result';
import Create from './pages/Create';

import {
  BrowserRouter as Router,
  Switch,
  Route
} from "react-router-dom";

function App() {
  return (
    <Router>
        <Switch>
          <Route exact path="/" component={Main}/>
          <Route exact path="/Admin_homepage" component={Admin_homepage}/>
          <Route exact path="/Admin_login" component={Admin_login}/>
          <Route path="/IC_inputted" component={IC_inputted}/>
          <Route path="/Event_result" component={Event_result}/>
          <Route exact path="/Create" component={Create}/>
        </Switch>
    </Router>
  );
}

export default App;
