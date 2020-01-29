import React, { Component } from 'react';
import Viewer from './Viewer';
import Updater from './Updater';
import Navigator from './Navigator';
import moment from 'moment';
import { ButtonToolbar, Button } from 'react-bootstrap';
import { Link } from 'react-router-dom';


import './ViewerTemplate.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import * as api from '../lib/api';

class ViewerTemplate extends Component {
    // 미세먼지 수치 값 state
    state = {
        // Set params from path.
        dt: this.props.match.params.dt,
        region: this.props.match.params.region,

        // update Tag
        isUpdate: false,

        // results state
        res : [
        {
            dt: '',
            region: '',
            no2: 0,
            o3: 0,
            co: 0,
            so2: 0,
            pm10: 0,
            pm25: 0,   
        },
        ]
        
    };

    // Server Call API
    // 1) Read
    readAPI = async (dt, region) => {
        try {
            const response = await api.readAPI(dt, region);
            const { dt: retrievedDT, region: retrievedRegion, no2, o3, co, so2, pm10, pm25} = response.data;

            this.setState({
                res : {dt:retrievedDT, region:retrievedRegion, no2, o3, co, so2, pm10, pm25}
            });

        } catch (e) {
            console.log(e)
            alert("HTTP Status : " + e.response.status + ",  " + e.response.data.message)
        }
    };

    // 2) Delete
    deleteAPI = async (dt, region) => {
        try {
            const response = await api.deleteAPI(dt, region);
            alert(response.data.message);
        } catch (e) {
            console.log(e)
        } finally {
            window.location.reload();
        }
    };

    // 3) Create
    createAPI = async (form) => {
        try {
            const response = await api.createAPI(form);
            alert(response.data.message);
        } catch (e) {
            console.log(e)
        } finally {
            window.location.reload();
        }
    };


    componentDidMount() {
        this.readAPI(this.state.dt, this.state.region);
    };



    // Prev / Next Handler
    handlePrev = () => {
        const prevDT = moment(this.state.dt, 'YYYYMMDD').subtract(1,'days').format('YYYYMMDD');

        this.setState({
            dt : prevDT,
        });

        const newPath = "/view/" + prevDT + "/" + this.state.region
        this.props.history.push(newPath)
        window.location.reload();
        // this.readAPI(this.state.dt, this.state.region);
    };

    handleNext = () => {
        const nextDT = moment(this.state.dt, 'YYYYMMDD').add(1,'days').format('YYYYMMDD');

        this.setState({
            dt : nextDT
        });

        const newPath = "/view/" + nextDT + "/" + this.state.region
        this.props.history.push(newPath)
        window.location.reload();
        // this.readAPI(this.state.dt, this.state.region);
    };

    handleCreate = (data) => {
        let form = new FormData()
        form.append('dt', data.dt)
        form.append('region',data.region)
        form.append('no2',data.no2)
        form.append('o3',data.o3)
        form.append('co',data.co)
        form.append('so2',data.so2)
        form.append('pm10',data.pm10)
        form.append('pm25',data.pm25)

        this.createAPI(form)

        this.setState({
            isUpdate: false,
        });
    }


    // Buttons for Add / Update / Remove
    onUpdate = () => {
        this.setState({
            isUpdate: true,
        })
    };


    onRemove = () => {
        this.deleteAPI(this.state.dt, this.state.region);
    };



    render () {
        const { dt, region, no2, o3, co, so2, pm10, pm25 } = this.state.res;

        if (this.state.isUpdate) {
            return (
                <div className="ViewerTemplate">
                    <div className="ViewerTemplate-Table">
                        <Updater 
                            onCreate={this.handleCreate}
                            dt={dt}
                            region={region}
                            no2={no2}
                            o3={o3}
                            co={co}
                            so2={so2}
                            pm10={pm10}
                            pm25={pm25}
                        />
                    </div>
                </div>
            );
        } 

        return (
            <div className="ViewerTemplate">
                <div>
                    <Navigator
                        onPrev={this.handlePrev}
                        onNext={this.handleNext}
                    />
                </div>
                <div className="ViewerTemplate-Table">
                    <Viewer 
                        dt={dt}
                        region={region}
                        no2={no2}
                        o3={o3}
                        co={co}
                        so2={so2}
                        pm10={pm10}
                        pm25={pm25}
                    />
                </div>
                <div className="ViewerTemplate-Button">
                    <ButtonToolbar>
                        <Link to ="/add">
                            <Button variant="primary">Add</Button>
                        </Link>
                        {dt !== undefined &&
                            <Button variant="warning" onClick={this.onUpdate}>Update</Button>
                        }                            
                        <Button variant="secondary" onClick={() => {if (window.confirm('Are you really sure to REMOVE this?')) this.onRemove() }}>Remove</Button>
                    </ButtonToolbar>
                </div>
            </div>
            );
    }
}


export default ViewerTemplate;