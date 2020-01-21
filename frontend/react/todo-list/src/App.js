import React, { Component } from 'react';
// import logo from './logo.svg';
import './App.css';
import TodoListTemplate from './component/TodoListTemplate';
import Form from './component/Form';

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

class App extends Component {
  render() {
    return (
      <div>
        <TodoListTemplate form={<Form/>} child="템플릿완성">
        </TodoListTemplate>
      </div>
    );
  }
}

export default App;
