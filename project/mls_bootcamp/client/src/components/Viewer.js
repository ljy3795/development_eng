import React from 'react';
import { Table } from 'react-bootstrap';
import moment from 'moment';

// Route in Route
const Viewer = ({dt, region, no2, o3, co, so2, pm10, pm25}) => {
    console.log(dt)
    return (
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
                    <td style={{width: "9.5vw"}}>{dt!==undefined && moment(dt,'YYYYMMDD').format("YYYY.MM.DD")}</td>
                    <td style={{width: "9.5vw"}}>{region}</td>
                    <td style={{width: "9.5vw"}}>{no2}</td>
                    <td style={{width: "9.5vw"}}>{o3}</td>
                    <td style={{width: "9.5vw"}}>{co}</td>
                    <td style={{width: "9.5vw"}}>{so2}</td>
                    <td style={{width: "9.5vw"}}>{pm10}</td>
                    <td style={{width: "9.5vw"}}>{pm25}</td>
                </tr>
            </tbody>
        </Table>
    );
};

export default Viewer;