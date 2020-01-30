import React, { Component } from 'react';
import Adder from './Adder';

import moment from 'moment';
import './AdderTemplate.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import * as api from '../lib/api';

class AdderTemplate extends Component {

    state = {
        dt : "YYYYMMDD",
        region : "지역",
    }

    handleCreate = async (data) => {
        try {
            // call readAPI by dt / region
            await api.readAPI(data.dt, data.region);
            // 이미 존재할 경우 confirm dialogue
            var r = window.confirm("Already the data exists, Do you want Update?");
        } catch (e) {
            // 존재하지 않을경우
            console.log(e)
        } finally {
            if(r === false){ 
                window.location.reload();
                return
            }

            let form = new FormData()

            form.append('dt', moment(data.dt).format("YYYYMMDD"))
            form.append('region',data.region)
            form.append('no2',data.no2)
            form.append('o3',data.o3)
            form.append('co',data.co)
            form.append('so2',data.so2)
            form.append('pm10',data.pm10)
            form.append('pm25',data.pm25)

            this.setState({
                dt: data.dt,
                region: data.region,
            })

            this.createAPI(form)
        }
        
        

    }

    // 3) Create API
    createAPI = async (form) => {
        try {
            const response = await api.createAPI(form);
            alert(response.data.message);
        } catch (e) {
            console.log(e)
        } finally {
            const newPath = "/view/" + this.state.dt + "/" + this.state.region
            this.props.history.push(newPath)
            // window.location.reload();
        }
    };

    render () {
        return (
            <div className="ViewerTemplate">
                <div className="ViewerTemplate-Table">
                    <Adder 
                        onCreate={this.handleCreate}
                    />
                </div>
            </div>
        );
    }
}



export default AdderTemplate;