import React, { Component } from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';

import ViewerTemplate from '../components/ViewerTemplate'
import AdderTemplate from '../components/AdderTemplate';
import HomeViewer from '../components/HomeViewer';

import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';


class App extends Component {
  render() {
    return (
        <div className='App'>
          {/* <div className='App-home'>
            <FaHome/>
          </div> */}
          <div className='App-header'>
            MLS-Bootcamp
          </div>
          <div className='App-footer'>
            (by Colin, 2020.1.27)
          </div>

          <BrowserRouter>
            <Switch>
              <Route exact path='/' component={HomeViewer}/>
              <Route exact path='/view/:dt/:region' component={ViewerTemplate}/>
              <Route exact path='/add/' component={AdderTemplate}/>
            </Switch>
          </BrowserRouter>

          {/* <Route
            path='/view/:dt/:region'
            component={(match) => <Viewer match={match}/>}
          /> */}
        </div>
    );
  }
}

export default App;