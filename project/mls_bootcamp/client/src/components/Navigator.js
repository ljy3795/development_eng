import React from 'react';
import {FaAngleLeft, FaAngleRight} from 'react-icons/fa';
import './Navigator.css';

const Navigator = ({ onPrev, onNext }) => {
    return (
        <div className='arrow'>
            <div className='left-end' onClick={onPrev}>
                <FaAngleLeft/>
            </div>
            <div className='right-end' onClick={onNext}>
                <FaAngleRight/>
            </div>
        </div>
    );
};

export default Navigator;