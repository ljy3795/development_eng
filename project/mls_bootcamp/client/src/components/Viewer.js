import React from 'react';
import { Table } from 'react-bootstrap';
import moment from 'moment';

// Route in Route
const Viewer = ({dt, region, no2, o3, co, so2, pm10, pm25}) => {
    return (
        <Table striped bordered hover style={{textAlign:'center', width:"100%"}}>
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
                    <td >{dt!==undefined && moment(dt,'YYYYMMDD').format("YYYY.MM.DD")}</td>
                    <td >{region}</td>
                    <td >{no2}</td>
                    <td >{o3}</td>
                    <td >{co}</td>
                    <td >{so2}</td>
                    <td >{pm10}</td>
                    <td >{pm25}</td>
                </tr>
            </tbody>
        </Table>
    );
};

export default Viewer;