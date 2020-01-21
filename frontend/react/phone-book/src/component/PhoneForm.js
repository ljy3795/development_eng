import React, { Component } from 'react';

class PhoneForm extends Component {
    state = {
        name : '',
        phone: ''
    }

    handleChange = (e) => {
        this.setState({
            // input의 name속성을 이용해서 한번에 처리 (handler를 여러개 사용하지 않아도 됨)
            [e.target.name]:e.target.value 
        })
    }

    handleSubmit = (e) => {
        // 페이지 리로딩 방지
        e.preventDefault(); // form에서는 submit을 하면 초기화 되기 때문에 초기화 방지
        // 상태값을 onCreate하여 부모에게 전달
        this.props.onCreate(this.state);
        // 상태초기화
        this.setState({
            name:'',
            phone:''
        })
    }

    //!!!!!!!!!!!!!!!!!!!!!!!!!!!
    // PhoneForm(App.js) -> onChange(PhoneForm.js) -> state(PhoneForm.js) -> onSubmit -> HandleSubmit -> onCreate (App.js -> 부모 component에 state를 전달) -> handleCreate
    //    onCreate는 parent Component에서 props로 넘겨줌
    render(){
        return (
            <form onSubmit={this.handleSubmit}>
                <input
                    placeholder='이름'
                    value={this.state.name}
                    // 값이 바뀔 때 마다 호출하는 event
                    onChange={this.handleChange}
                    name="name"
                />
                <input
                    placeholder='전화번호'
                    value={this.state.phone}
                    onChange={this.handleChange}
                    name="phone"
                />
                <button type="submit">등록</button>
                <div>{this.state.name} {this.state.phone}</div>
            </form>
        );
    }
}

export default PhoneForm;