// import React from 'react';
// import logo from './logo.svg';
// import './App.css';

// function App() {
//   return (
//     <div className="App">
//       <header className="App-header">
//         <img src={logo} className="App-logo" alt="logo" />
//         <p>
//           Edit <code>src/App.js</code> and save to reload.
//         </p>
//         <a
//           className="App-link"
//           href="https://reactjs.org"
//           target="_blank"
//           rel="noopener noreferrer"
//         >
//           Learn React
//         </a>
//       </header>
//     </div>
//   );
// }

// export default App;



import React, { Component } from 'react';
import './App.css'
// import MyName from './MyName'
// import Counter from './Counter'
// import CounterAPI from './CounterAPI'
import CounterAPIError from './CounterAPIError'

class App extends Component {
  render() {
    // const style = {
    //   backgroundColor: 'black',
    //   padding: '16px',
    //   color: 'blue',
    //   fontSize:'20px'
    // };

    return (
      // <MyName name="하이"/>
      <CounterAPIError />
    );
  }
}

export default App;
