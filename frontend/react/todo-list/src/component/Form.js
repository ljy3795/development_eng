import React from 'react';
import './Form.css';

const Form = ({value, onChange, onCreate, onKeyPress}) => {
    return (
        <div className="form">
            {/* App.js -> TodoListTemplate -> Form.js -> onChange(App.js의 props method) -> App.js의 handleChange로 전달 */}
            <input value={value} onChange={onChange} onKeyPress={onKeyPress}></input>
            <div className="create-button" onClick={onCreate}>
                추가
            </div>
        </div>
    )
}

export default Form