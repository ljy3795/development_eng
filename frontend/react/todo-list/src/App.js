import React, { Component } from 'react';
// import logo from './logo.svg';
import './App.css';
import TodoListTemplate from './component/TodoListTemplate';
import Form from './component/Form';
import TodoItemList from './component/TodoItemList';


class App extends Component {
  id = 3;

  state = {
    input : '',
    todos: [
      { id : 0, text: ' 리액트 소개', checked: false},
      { id : 1, text: ' 리액트 소개', checked: true},
      { id : 2, text: ' 리액트 소개', checked: false}
    ]
  }

  handleChange = (e) => {
    this.setState({
      input: e.target.value // input의 다음 바뀔 값
    });
  }

  handleCreate = () => {
    const { input, todos } = this.state;
    this.setState({
      input: '', // input 초기화
      todos: todos.concat({
        id: this.id++,
        text: input,
        checked:false
      })
    });
  }

  handleKeyPress = (e) => {
    // Enter입력 시 handleCreate 호출
    if(e.key === 'Enter'){
      this.handleCreate();
    }
  }

  handleToggle = (id) => {
    const { todos } = this.state;

    // id를 가지고 몇번째 아이템인지 찾음
    const index = todos.findIndex(todo => todo.id === id);
    const selected = todos[index];

    // copy
    const nextTodos = [...todos];

    // 기존값들을 복사하고, 선택된 아이템만 값을 변경
    nextTodos[index] = {
      ...selected,
      checked:!selected.checked
    }

    this.setState({
      todos:nextTodos
    })
  }


  handleRemove = (id) => {
    const { todos } = this.state;
    this.setState({
      todos : todos.filter(info => info.id !== id)
    })
  }


  render() {
    const {input, todos} = this.state;

    return (
      <div>
        <TodoListTemplate 
          form={<Form
                  value={input}
                  onKeyPress={this.handleKeyPress}
                  onChange={this.handleChange}
                  onCreate={this.handleCreate}
          />} 
          child={<TodoItemList
                    todos={todos}
                    onToggle={this.handleToggle}
                    onRemove={this.handleRemove}
                />}>
        </TodoListTemplate>
      </div>
    );
  }
}

export default App;
