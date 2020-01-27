import React, { Component } from 'react';
import Adder from './Adder';


import './AdderTemplate.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import * as api from '../lib/api';

class AdderTemplate extends Component {

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
    }

    // 3) Create API
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