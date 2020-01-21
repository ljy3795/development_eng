import React, { Component } from 'react';
import './TodoItem.css'

class TodoItem extends Component {

    handleToggle = () => {
        const { id, onToggle } = this.props;
        onToggle(id);
    }

    handleRemove = (e) => {
        const {id, onRemove} = this.props;
        e.stopPropagation(); // onToggle이 실행되지 않도록 함 (상위 component까지 안가도록 event 확산을 방지)
        onRemove(id);
    }

    shouldComponentUpdate(nextProps, nextState) {
        return this.props.checked !== nextProps.checked;
    }


    render() {
        const { text, checked, id} = this.props;
        console.log(id);

        return (
            <div className="todo-item" onClick={this.handleToggle}>
                <div className="remove" onClick={this.handleRemove}
                >&times;
                </div>

                <div className={`todo-text ${checked && 'checked'}`}>
                    <div>{text}</div>
                </div>
                {
                    checked && (<div className="check-mark"></div>)
                }
            </div>
        )
    }
}

export default TodoItem