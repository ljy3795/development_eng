import React, { Component } from 'react';
import TodoItem from './TodoItem';

class TodoItemList extends Component {
    static defaultProps = {
        todos: []
    }

    shouldComponentUpdate(nextProps, nextState) {
        return this.props.todos !== nextProps.todos;
    }

    render() {
        const {todos, onToggle, onRemove} = this.props;
        const list = todos.map(
            ({id, text, checked}) => (
                    <TodoItem
                        id={id}
                        text={text}
                        checked={checked}
                        onToggle={onToggle}
                        onRemove={onRemove}
                        >
                    </TodoItem>)
        )

        return (
            <div>
                {list}
            </div>
        );
    }
}

export default TodoItemList;