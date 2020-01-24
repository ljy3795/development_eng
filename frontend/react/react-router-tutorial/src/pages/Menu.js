import React from 'react';
import { Link, NavLink } from 'react-router-dom';

const Menu = () => {
    const activeStyle = {
        color: 'green',
        fontSize:'2rem'
    };

    return (
        <div>
            <ul>
                {/* Link를 이용해서 해당 주소로 감 */}
                {/* <li><Link to="/">Home</Link></li>
                <li><Link to="/about">About</Link></li>
                <li><Link to="/about/foo?detail=true">About Foo</Link></li> */}
                <li><NavLink exact to="/" activeStyle={activeStyle}>Home</NavLink></li>
                <li><NavLink exact to="/about" activeStyle={activeStyle}>About</NavLink></li>
                <li><NavLink to="/about/foo?detail=true" activeStyle={activeStyle}>About Foo</NavLink></li>
                <li><NavLink to="/posts" activeStyle={activeStyle}>Posts</NavLink></li>
            </ul>
        </div>
    );
};

export default Menu;