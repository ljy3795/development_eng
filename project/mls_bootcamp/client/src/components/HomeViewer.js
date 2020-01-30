import React, { Component } from 'react';
import { Form, Button } from 'react-bootstrap';

import 'bootstrap/dist/css/bootstrap.min.css';

class HomeViewer extends Component {

    state = {
        dt: "YYYYMMDD",
        region: "지역",
        
        dtValid: false,
    };

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

    handleView = (e) => {
        // 페이지 리로딩 방지
        e.preventDefault(); // form에서는 submit을 하면 초기화 되기 때문에 초기화 방지
        if (this.state.dtValid){
            const { dt, region } = this.state
            // 상태초기화
            this.setState({
                dt : "YYYYMMDD",
                region : "지역",

                dtValid : false,
            })

            const newPath = "/view/" + dt + "/" + region
            this.props.history.push(newPath)
            // window.location.reload();

        } else {
            alert("Date type should be 'YYYYMMDD' in number")
        }
    }

    handleAdd = (e) => {
        // 페이지 리로딩 방지
        e.preventDefault(); // form에서는 submit을 하면 초기화 되기 때문에 초기화 방지
        // 상태초기화
        this.setState({
            dt : "YYYYMMDD",
            region : "지역",

            dtValid : false,
        })
        this.props.history.push("/add")
    }

    
    render () {
        return (
            <div style={{width:'30%', position:'absolute', left:"30%", paddingTop:"3vh"}}>
                <Form>
                    <Form.Group controlId="formDate">
                        <Form.Label style={{fontWeight:'bold', fontSize:'150%'}}>Date</Form.Label>
                        <Form.Control type="string" placeholder="YYYYMMDD" name="dt" onChange={this.handleChange} required/>
                        <Form.Text className="text-muted">
                        (Input should be in YYYYMMDD format)
                        </Form.Text>
                    </Form.Group>

                    <Form.Group controlId="formRegion">
                        <Form.Label style={{fontWeight:'bold', fontSize:'150%'}}>Region</Form.Label>
                        <Form.Control type="string" placeholder="지역" name="region" onChange={this.handleChange} required/>
                    </Form.Group>
                    <Button style={{float:'left'}} variant="primary" type="submit" onClick={this.handleView}>
                        View
                    </Button>
                    <Button style={{float:'right'}} variant="warning" type="submit" onClick={this.handleAdd}>
                        Add
                    </Button>
                </Form>
            </div>
        )
    }
}


export default HomeViewer;