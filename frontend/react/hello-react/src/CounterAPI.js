import React, { Component } from 'react';

class CounterAPI extends Component {
  state = {
    number: 0
  }

  // 컴포넌트 초기
  constructor(props) {
    super(props);
    console.log('constructor');
  }
  
  componentWillMount() {
    console.log('componentWillMount (deprecated)');
  }

  // 컴포넌트가 나간 이후에 호출 (render 이후)
  componentDidMount() {
    console.log('componentDidMount');
  }



  // 컴포넌트 업데이트
  shouldComponentUpdate(nextProps, nextState) {
    // 5 의 배수라면 리렌더링 하지 않음
    console.log('shouldComponentUpdate 1)');
    if (nextState.number % 5 === 0) return false;
    return true;
  }

  // shouldComponentUpdate가 true일 때 만 호출
  componentWillUpdate(nextProps, nextState) {
    console.log('componentWillUpdate 2)');
  }
  
  componentDidUpdate(prevProps, prevState) {
    console.log('componentDidUpdate 3)');
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
  
  render() {
    console.log('render');
    return (
      <div>
        <h1>카운터</h1>
        <div>값: {this.state.number}</div>
        <button onClick={this.handleIncrease}>+</button>
        <button onClick={this.handleDecrease}>-</button>
      </div>
    );
  }
}

export default CounterAPI;