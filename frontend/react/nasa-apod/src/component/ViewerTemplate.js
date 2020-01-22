import React from 'react';
import './ViewerTemplate.css';


// JSX형태의 props인 viewer와 spaceNavigator를 받아와서 적당한 위치에 렌더링
const ViewerTemplate = ({viewer, spaceNaviagor}) => {
    return (
        <div className="ViewerTemplate">
            <header className="ViewerTemplate-header">
                Astronomy Picture of the Day
            </header>
            <div>
                {viewer}
            </div>
            <div>
                {spaceNaviagor}
            </div>
        </div>
    )
}

export default ViewerTemplate;