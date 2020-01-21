import React, { Component } from 'react';

class PhoneInfo extends Component {
    static defaultProps = {
        info : {
            name : '이름',
            phone : '010-0000-0000',
            id:0
        }
    }

    state = {
        // 우리는 수정 버튼을 눌렀을 떄 editing 값을 true 로 설정해줄것입니다.
        // 이 값이 true 일 때에는, 기존에 텍스트 형태로 보여주던 값들을
        // input 형태로 보여주게 됩니다.
        editing: false,
        // input 의 값은 유동적이겠지요? input 값을 담기 위해서 각 필드를 위한 값도
        // 설정합니다
        name: '',
        phone: '',
    }

    handleRemove = () => {
        // 삭제 버튼이 클릭되면 onRemove에 id를 넣어서 호출
        const {info, onRemove} = this.props;
        onRemove(info.id); // PhoneInforLIst -> App.js를 타고가서 최종 parent의 onRemove를 실행
    }

    // editing 값을 반전시키는 함수입니다
    // true -> false, false -> true
    handleToggleEdit = () => {
        const { editing } = this.state;
        this.setState({ 
            editing: !editing 
        });
    }

      // input 에서 onChange 이벤트가 발생 될 때
    // 호출되는 함수입니다
    handleChange = (e) => {
        // 각 handleChange의 name 및 value
        const { name, value } = e.target;
        // console.log(name, value)
        this.setState({
        [name]: value
        });
        // console.log(this.state.name, this.state.phone)
    }


    // PhoneInfo의 state를 이용하여 수정할 때의 JSX의 변수를 전달
    componentDidUpdate(prevProps, prevState) {
        // 여기서는, editing 값이 바뀔 때 처리 할 로직이 적혀있습니다.
        // 수정을 눌렀을땐, 기존의 값이 input에 나타나고,
        // 수정을 적용할땐, input 의 값들을 부모한테 전달해줍니다.
    
        const { info, onUpdate } = this.props;
        if(!prevState.editing && this.state.editing) {
          // editing 값이 false -> true 로 전환 될 때
          // info 의 값을 state 에 넣어준다

          console.log(1)
          console.log("ID ", info.id)
          console.log(info.name)
          console.log(info.phone)
          this.setState({
            name: info.name,
            phone: info.phone
          })
        }
    
        // 이 시점에서는 이미 render() 이후 바뀐 state임
        if (prevState.editing && !this.state.editing) {
            console.log(2)
            console.log(info.id)
            console.log(this.state.name)
            console.log(this.state.phone)
          // editing 값이 true -> false 로 전환 될 때
          onUpdate(info.id, {
            name: this.state.name,
            phone: this.state.phone
          });
        }
    }


    // 수정상태가 아닌 component (array 값들)을 아끼기 위해 rendering을 하지 않음
    // 즉, DOM변화가 일어나지 않지만 Virtual DOM에 그리는 자원도 아껴주기 위해 사용
    shouldComponentUpdate(nextProps, nextState) {
        // 수정상태가 아니고, info값이 같다면 리렌더링 안함
        //   즉, 아래 render()를 하지 않음
        if (!this.state.editing && !nextState.editing && nextProps.info === this.props.info){
            return false;
        }

        return true;
    }

    render() {
        console.log('render PhoneInfo ' + this.props.info.id)
        const style = {
            border: '1px solid black',
            padding: '8px',
            margin: '8ps'
        };

        const { editing } = this.state;

        // 수정모드
        if (editing) {
            return (
                <div stype={style}>
                    <div>
                        <input
                            value={this.state.name}
                            name="name"
                            placeholder="이름"
                            onChange={this.handleChange}    
                        >
                        </input>
                    </div>
                    <div>
                        <input
                            value={this.state.phone}
                            name="phone"
                            placeholder="전화번호"
                            onChange={this.handleChange}    
                        >
                        </input>
                    </div>
                    <button onClick={this.handleToggleEdit}>적용</button>
                    <button onClick={this.handleRemove}>삭제</button>
                </div>
            );
        }


        // 일반모드
        const {
            name, phone
        } = this.props.info;

        return (
            <div style={style}>
                <div><b>{name}</b></div>
                <div>{phone}</div>
                {/* <div>aa</div> */}
                <button onClick={this.handleToggleEdit}>수정</button>
                <button onClick={this.handleRemove}>삭제</button>
            </div>
        );
    }
}

export default PhoneInfo;