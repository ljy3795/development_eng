import React, { Component } from 'react';
import { Table, Button } from 'react-bootstrap';
import './Adder.css';


// Route in Route
class Updater extends Component {
    state = {
        dt : this.props.dt,
        region : this.props.region,
        no2 : this.props.no2,
        o3 : this.props.o3,
        co : this.props.co,
        so2 : this.props.so2,
        pm10 : this.props.pm10,
        pm25 : this.props.pm25,
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
    }

    render(){
        return(
            <form onSubmit={this.handleSubmit}>
                <Table striped bordered hover style={{'text-align':'center'}}>
                    <thead>
                        <tr>
                        <th>Date</th>
                        <th>Region</th>
                        <th>NO2</th>
                        <th>O3</th>
                        <th>CO</th>
                        <th>SO2</th>
                        <th>PM10</th>
                        <th>PM25</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <td>
                                {this.state.dt}
                            </td>
                            <td>
                                {this.state.region}
                            </td>
                            <td>
                                <input  type="number"
                                        name="no2"
                                        value={this.state.no2}
                                        onChange={this.handleChange}
                                        style={{width: "9.5vw"}}
                                        required/>
                            </td>
                            <td>
                                <input  type="number"
                                        name="o3"
                                        value={this.state.o3}
                                        onChange={this.handleChange}
                                        style={{width: "9.5vw"}}
                                        required/>
                            </td>
                            <td>
                                <input  type="number"
                                        name="co"
                                        value={this.state.co}
                                        onChange={this.handleChange}
                                        style={{width: "9.5vw"}}
                                        required/>
                            </td>
                            <td>
                                <input  type="number"
                                        name="so2"
                                        value={this.state.so2}
                                        onChange={this.handleChange}
                                        style={{width: "9.5vw"}}
                                        required/>
                            </td>
                            <td>
                                <input  type="number"
                                        name="pm10"
                                        value={this.state.pm10}
                                        onChange={this.handleChange}
                                        style={{width: "9.5vw"}}
                                        required/>
                            </td>
                            <td>
                                <input  type="number"
                                        name="pm25"
                                        value={this.state.pm25}
                                        onChange={this.handleChange}
                                        style={{width: "9.5vw"}}
                                        required/>
                            </td>
                        </tr>
                    </tbody>
                </Table>
                <div className='Adder-header'>
                    <Button variant="primary" type="submit">
                        업데이트
                    </Button>
                </div>
            </form>
        );
    }
}

export default Updater;