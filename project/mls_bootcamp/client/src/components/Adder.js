import React, { Component } from 'react';
import { Table, Button } from 'react-bootstrap';
import './Adder.css';





// Route in Route
class Viewer extends Component {
    static defaultProps = {
        dt : '',
        region : '',
        no2 : '',
        o3 : '',
        co : '',
        so2 : '',
        pm10 : '',
        pm25 : '',
    }

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
        // 상태초기화
        this.setState({
            dt : '',
            region : '',
            no2 : '',
            o3 : '',
            co : '',
            so : '',
            pm10 : '',
            pm25 : '',
        })
    }

    render(){
        // console.log(this.state)
        return(
            <form onSubmit={this.handleSubmit}>
                <Table striped bordered hover>
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
                                <input  type="text"
                                        name="dt"
                                        value={this.state.dt}
                                        onChange={this.handleChange}
                                        style={{width: "9.5vw"}}/>
                            </td>
                            <td>
                                <input  type="text"
                                        name="region"
                                        value={this.state.region}
                                        onChange={this.handleChange}
                                        style={{width: "9.5vw"}}/>
                            </td>
                            <td>
                                <input  type="text"
                                        name="no2"
                                        value={this.state.no2}
                                        onChange={this.handleChange}
                                        style={{width: "9.5vw"}}/>
                            </td>
                            <td>
                                <input  type="text"
                                        name="o3"
                                        value={this.state.o3}
                                        onChange={this.handleChange}
                                        style={{width: "9.5vw"}}/>
                            </td>
                            <td>
                                <input  type="text"
                                        name="co"
                                        value={this.state.co}
                                        onChange={this.handleChange}
                                        style={{width: "9.5vw"}}/>
                            </td>
                            <td>
                                <input  type="text"
                                        name="so2"
                                        value={this.state.so2}
                                        onChange={this.handleChange}
                                        style={{width: "9.5vw"}}/>
                            </td>
                            <td>
                                <input  type="text"
                                        name="pm10"
                                        value={this.state.pm10}
                                        onChange={this.handleChange}
                                        style={{width: "9.5vw"}}/>
                            </td>
                            <td>
                                <input  type="text"
                                        name="pm25"
                                        value={this.state.pm25}
                                        onChange={this.handleChange}
                                        style={{width: "9.5vw"}}/>
                            </td>
                        </tr>
                    </tbody>
                </Table>
                <div className='Adder-header'>
                    <Button variant="primary" type="submit">
                        등록
                    </Button>
                </div>
            </form>
        );
    }
}

export default Viewer;