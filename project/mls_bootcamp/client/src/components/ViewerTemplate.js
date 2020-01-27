import React, { Component } from 'react';
import Viewer from './Viewer';
import Adder from './Adder';
import Navigator from './Navigator';
import moment from 'moment';
import { ButtonToolbar, Button } from 'react-bootstrap';


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
            no2: '',
            o3: '',
            co: '',
            so2: '',
            pm10: '',
            pm25: '',   
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

    componentDidMount() {
        this.readAPI(this.state.dt, this.state.region);
    };

    handlePrev = () => {
        const dt = this.state.dt;
        const prevDT = moment(dt, 'YYYYMMDD').subtract(1,'days').format('YYYYMMDD');

        this.setState({
            dt : prevDT
        });
        const newPath = "/view/" + this.state.dt + "/" + this.state.region
        this.props.history.push(newPath)
        this.readAPI(this.state.dt, this.state.region);
        // window.location.reload();
    };

    handleNext = () => {
        const dt = this.state.dt;
        const nextDT = moment(dt, 'YYYYMMDD').add(1,'days').format('YYYYMMDD');
        // <Redirect to='/view/20191231/강남구'/>

        this.setState({
            dt : nextDT
        });
        const newPath = "/view/" + this.state.dt + "/" + this.state.region
        this.props.history.push(newPath)
        this.readAPI(this.state.dt, this.state.region);
        // window.location.reload();
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

        console.log(this.state.isUpdate)
        this.setState({
            isUpdate: false,
        }, function () {
            console.log(this.state.isUpdate);
        });
    }

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

    onAdd = () => {
        // <Route path='/add' component={Adder}/>
        this.props.history.push("/add")
        window.location.reload();

        alert('Go to ADD Page!');
    };

    onUpdate = () => {
        this.setState({
            isUpdate: true,
        })
        alert('Go to Update Page!');
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
                        <Adder 
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
        } else {
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
                            <Button variant="primary" onClick={this.onAdd}>Add</Button>
                            <Button variant="warning" onClick={this.onUpdate}>Update</Button>
                            <Button variant="secondary" onClick={this.onRemove}>Remove</Button>
                        </ButtonToolbar>
                    </div>
                </div>
            );
        }
    }
}



export default ViewerTemplate;