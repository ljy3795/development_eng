import React, { Component } from 'react';
import { Table, Button } from 'react-bootstrap';
import './Adder.css';


// Route in Route
class Adder extends Component {
    state = {
        dt : "YYYYMMDD",
        region : "지역",
        no2 : 0,
        o3 : 0,
        co : 0,
        so2 : 0,
        pm10 : 0,
        pm25 : 0,

        dtValid: false,
    }

    handleChange = (e) => {
        // YYYYMMDD(number) regexp
        const dtTypeRegExp = /^\d{8}$/;

        // input이 dt인 경우, regexp 확인
        if (e.target.name === "dt" && e.target.value.match(dtTypeRegExp)){
            this.setState({
                dtValid: true,
            })
        } 

        this.setState({
            // input의 name속성을 이용해서 한번에 처리 (handler를 여러개 사용하지 않아도 됨)
            [e.target.name]:e.target.value 
        })
    }

    handleSubmit = (e) => {
        // 페이지 리로딩 방지
        e.preventDefault(); // form에서는 submit을 하면 초기화 되기 때문에 초기화 방지
        // 상태값을 onCreate하여 부모에게 전달. 단 dt값이 valid할 경우
        if (this.state.dtValid) {
            this.props.onCreate(this.state);
            // 상태초기화
            this.setState({
                dt : "YYYYMMDD",
                region : "지역",
                no2 : 0,
                o3 : 0,
                co : 0,
                so2 : 0,
                pm10 : 0,
                pm25 : 0,

                dtValid : false,
            })
        } else {
            alert("Date type should be 'YYYYMMDD' in number")
        }
    }

    render(){
        return(
            <form onSubmit={this.handleSubmit}>
                <Table striped bordered hover style={{textAlign:'center', width:'100%'}}>
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
                                        placeholder={this.state.dt}
                                        onChange={this.handleChange}
                                        required
                                        style={{width:'100%'}}/>
                            </td>
                            <td>
                                <input  type="text"
                                        name="region"
                                        placeholder={this.state.region}
                                        onChange={this.handleChange}
                                        required
                                        style={{width:'100%'}}/>
                            </td>
                            <td>
                                <input  type="number"
                                        name="no2"
                                        step="0.001"
                                        placeholder={this.state.no2}
                                        onChange={this.handleChange}
                                        required
                                        style={{width:'100%'}}/>
                            </td>
                            <td>
                                <input  type="number"
                                        name="o3"
                                        step="0.001"
                                        placeholder={this.state.o3}
                                        onChange={this.handleChange}
                                        required                                        
                                        style={{width:'100%'}}/>
                            </td>
                            <td>
                                <input  type="number"
                                        name="co"
                                        step="0.001"
                                        placeholder={this.state.co}
                                        onChange={this.handleChange}
                                        required
                                        style={{width:'100%'}}/>
                            </td>
                            <td>
                                <input  type="number"
                                        name="so2"
                                        step="0.001"
                                        placeholder={this.state.so2}
                                        onChange={this.handleChange}
                                        required
                                        style={{width:'100%'}}/>
                            </td>
                            <td>
                                <input  type="number"
                                        name="pm10"
                                        step="0.001"
                                        placeholder={this.state.pm10}
                                        onChange={this.handleChange}
                                        required
                                        style={{width:'100%'}}/>
                            </td>
                            <td>
                                <input  type="number"
                                        name="pm25"
                                        step="0.001"
                                        placeholder={this.state.pm25}
                                        onChange={this.handleChange}
                                        required
                                        style={{width:'100%'}}/>
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

export default Adder;