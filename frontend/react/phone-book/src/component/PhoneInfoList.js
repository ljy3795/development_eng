import React, { Component } from 'react';
import PhoneInfo from './PhoneInfo'

class PhoneInfoList extends Component {
    static defaultProps = {
        data: []
    }

    render() {
        // data라는 배열을 가져와서 JSX로 변환해줌        
        const { data } = this.props;

        const list = data.map(
            info => (<PhoneInfo key={info.id} info={info}></PhoneInfo>)
            // 배열의 key값은 필수
            // data에 존재하는 value값을 info로 보고 계산
        );
        console.log(list)


        return (
            <div>
                {list}
            </div>
        )

    }
}

export default PhoneInfoList;