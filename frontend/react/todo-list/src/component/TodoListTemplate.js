import React from 'react';
import './TodoListTemplate.css'

// form, child는 parent로 받은 props
const TodoListTemplate = ({form, child}) => {
    return (
        <main className="todo-list-template">
            <div className="title">
                오늘 할 일
            </div>
            <section className="form-wrapper">
                {form}
            </section>
            <section className="todos-wrapper">
                {child}    
            </section>
        </main>

    );
};

export default TodoListTemplate;