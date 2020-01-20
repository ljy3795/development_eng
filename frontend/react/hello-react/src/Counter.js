import React, { Component } from 'react';

class Counter extends Component {
    // state variable
    state = {
        number : 0
    }

    // method
    // HandleIncrease = () => {
    //     this.setState({
    //         number: this.state.number + 1
    //     });
    // }
    // HandleIncrease = () => {
    //     this.setState(
    //         (state) => ({
    //             number: state.number + 1
    //         })
    //     );
    // }

    HandleIncrease = () => {
        this.setState(
            ({number}) => ({
                number: number + 1
            })
        );
    }

    // method
    // HandleDecrease = () => {
    //     this.setState({
    //         number: this.state.number - 1
    //     });
    // }

    HandleDecrease = () => {
        // 비구조화 할당
        const {number} = this.state;
        this.setState({
            number: number -1
        });
    }


    render() {
        return (
            <div>
                <h1>카운터</h1>
                <div>값: {this.state.number}</div>
                <button onClick={this.HandleIncrease}>+</button>
                <button onClick={this.HandleDecrease}>-</button>
            </div>
        )
    }
}

export default Counter;

