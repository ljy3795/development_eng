import React from 'react';
import './Viewer.css';
import {
    ChasingDots
} from 'better-react-spinkit'

// 이미지 혹은 동영상을 보여주는 것
const Viewer = ({mediaType, url, loading}) => {
    if (!url) return null;

    if(loading) {
        // 로딩중일때 로더 보여주기
        return  <div className='Viewer'>
                    {console.log('loading')}
                    <ChasingDots color="white" size={60}/>
                </div>
    }

    return (
        <div className='Viwer'>
            {
                mediaType === 'image' ? (
                    <img className='viewerImage' onClick={() => window.open(url)} src={url} alt="space"/>
                ) : (
                    <iframe className='viewerVideo' title="space-video" src={url} frameBorder="0" allow="autoplay" allowFullScreen></iframe>
                )
            }
        </div>
    );
};

export default Viewer;