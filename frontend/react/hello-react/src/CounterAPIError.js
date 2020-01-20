import React, { Component } from 'react';

// 컴포넌트
const Problematic = () => {
  throw (new Error('버그가 나타났다!'));
  return (
    <div>
      
    </div>
  );
};

class CounterAPIError extends Component {
    // state variable
    state = {
        number : 0,
        error : false
    }
    handleIncrease = () => {
        const { number } = this.state;
        this.setState({
          number: number + 1
        });
      }
    
    handleDecrease = () => {
    this.setState(
        ({ number }) => ({
        number: number - 1
        })
    );
    }

    componentDidCatch(error, info){
        console.log("AA")
        this.setState({
            error:true
        });
    }
  
    render() {
        return (
        <div>
            <h1>카운터</h1>
            <div>값: {this.state.number}</div>
            { this.state.number === 4 && <Problematic /> }
            <button onClick={this.handleIncrease}>+</button>
            <button onClick={this.handleDecrease}>-</button>
        </div>
        );
    }
}

export default CounterAPIError;