import React, { Component } from 'react';


// Class 방식
// class MyName extends Component {
//     // static defaultProps = {
//     //     name : "기본이름"
//     // }
//     render() {
//       return (
//         <div>
//     안녕하세요! 제 이름은 <b>{this.props.name}</b> 입니다.
//         </div>
//       )
//     }
// }


// 함수 방식
const MyName = ({name}) => {
    return (
        <div>
            안녕하세요! <b>{name}</b>
        </div>
    )
}


MyName.defaultProps = {
    name : "헬로우"
}

export default MyName;