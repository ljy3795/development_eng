import React, { Component } from 'react';
import { Route, Switch } from 'react-router-dom';
import { Home, About, Menu, Posts } from 'pages';

class App extends Component {
    render() {
        return (
            <div>
                <div>
                    <Menu/>
                </div>
                <div>
                    <Route exact path="/" component={Home}/>
                    {/* 첫번 째 매칭되는 component만 보여주고 나머지는 안보여줌 */}
                    <Switch>
                        <Route path="/about/:name" component={About}/>
                        <Route path="/about" component={About}/>
                    </Switch>
                    <Route path="/posts" component={Posts}/>
                </div>  
            </div>
        );
    }
}

export default App;