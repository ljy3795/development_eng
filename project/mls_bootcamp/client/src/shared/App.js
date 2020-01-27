import React, { Component } from 'react';
import { Route } from 'react-router-dom';

import ViewerTemplate from '../components/ViewerTemplate'
import AdderTemplate from '../components/AdderTemplate';

import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';


class App extends Component {
  render() {
    return (
        <div className='App'>
          <div className='App-header'>
            MLS-Bootcamp
          </div>
          <div className='App-footer'>
            (by Colin, 2020.1.27)
          </div>

          <Route path='/view/:dt/:region' component={ViewerTemplate}/>
          <Route path='/add/' component={AdderTemplate}/>

          {/* <Route
            path='/view/:dt/:region'
            component={(match) => <Viewer match={match}/>}
          /> */}
        </div>
    );
  }
}

export default App;
