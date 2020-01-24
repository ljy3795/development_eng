
// 일반적인 방법
// import React, { Component } from 'react';
// import notify from './notify';

// class App extends Component {
//   handleClick = () => {
//     notify();
//   };

//   render() {
//     return (
//       <div>
//         <button onClick={this.handleClick}>Click Me!</button>
//       </div>
//     );
//   }
// }

// export default App;


// code splitting

import React, { Component } from 'react';

class App extends Component {
  handleClick = () => {
    import('./notify').then(({default: notify}) => {
        notify();
    });
  };

  render() {
    return (
      <div>
        <button onClick={this.handleClick}>Click Me!</button>
      </div>
    );
  } 
}

export default App;