import React, { Component } from 'react';
import PhoneInfo from './PhoneInfo'

class PhoneInfoList extends Component {
    static defaultProps = {
        data: [],
        onRemove: () => console.warn('onRemove not defined'),
        onUpdate: () => console.warn('onUpdate not defined')
    }

    render() {
        // data라는 배열을 가져와서 JSX로 변환해줌        
        // App.js의 state를 가져옴 (information)
        const { data, onRemove, onUpdate } = this.props;

        // App.js의 information의 값을 PhoneInfo를 통해서 array로 return
        // PhoneInfoList(App.js) -> data -> PhoneInfo (PhoneInfoList의 child) -> JSX return -> array  
        const list = data.map(
            info => (<PhoneInfo 
                        key={info.id} 
                        info={info}
                        // App.js에서 받은 onRemove를 props로 받아서 넘김
                        onRemove={onRemove}
                        onUpdate={onUpdate}>
                    </PhoneInfo>)
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