import React, {useRef, useState} from 'react';
import {Link} from "react-router-dom";
import {
    ArrowDown,
    Friends,
    Gaming,
    HomeActive,
    Logo,
    Market,
    Menu,
    Messenger,
    Notifications,
    Search,
    Watch
} from "../../svg";
import './style.css';
import SearchMenu from "./SearchMenu";
import AllMenu from "./AllMenu";
import UserMenu from "./UserMenu/UserMenu";
import {useClickOutside} from "../../hooks/useClickOutside";
import classNames from "classnames";

const color = '#65676b';

function Header({user}) {
    const [showSearchMenu, setShowSearchMenu] = useState(false);
    const [showAll, setShowAll] = useState(false);
    const [showUserMenu, setShowUserMenu] = useState(false);
    const userMenuRef = useRef(null)

    useClickOutside(userMenuRef, () => {
        setShowUserMenu(false);
    })
    return (
        <header>
            <div className="header_left">
                <Link to="/" className='header_logo'>
                    <div className="circle">
                        <Logo/>
                    </div>
                </Link>
                <div className="search search1" onClick={() => setShowSearchMenu(true)}>
                    <Search color={color}/>
                    <input type="text"
                           placeholder='Search Facebook'
                           className='hide_input'/>
                </div>
            </div>
            {
                showSearchMenu && <SearchMenu
                    setShowSearchMenu={setShowSearchMenu}
                    color={color}/>
            }
            <div className="header_middle">
                <Link className='middle_icon active' to='/'>
                    <HomeActive/>
                </Link>
                <Link className='middle_icon hover1' to='/'>
                    <Friends color={color}/>
                </Link>
                <Link className='middle_icon hover1' to='/'>
                    <Watch color={color}/>
                    <div className="middle_notification">8+</div>
                </Link>
                <Link className='middle_icon hover1' to='/'>
                    <Market color={color}/>
                </Link>
                <Link className='middle_icon hover1' to='/'>
                    <Gaming color={color}/>
                </Link>
            </div>
            <div className="header_right">
                <Link to='/profile' className="profile_link hover1">
                    <img src={user?.picture} alt={user?.user_name}/>
                    <span>{user?.first_name}</span>
                </Link>
                <div className={classNames("circle_icon hover1", {
                    'active_header': showAll
                })} onClick={() => setShowAll(prev => !prev)}>
                    <Menu/>
                    {showAll && <AllMenu setShowAll={setShowAll}/>}
                </div>
                <div className="circle_icon hover1">
                    <Messenger/>
                </div>
                <div className="circle_icon hover1">
                    <Notifications/>
                    <div className="right_notification">7</div>
                </div>
                <div className={classNames("circle_icon hover1", {
                    'active_header': showUserMenu
                })} ref={userMenuRef} onClick={() => setShowUserMenu(true)}>
                    <ArrowDown/>
                    {showUserMenu && <UserMenu user={user}/>}
                </div>
            </div>
        </header>
    );
}

export default Header;