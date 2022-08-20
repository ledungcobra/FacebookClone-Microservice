import React, {useEffect, useRef, useState} from 'react';
import {Return, Search} from "../../svg";
import {useClickOutside} from "../../hooks/useClickOutside";
import {ref} from "yup";

function SearchMenu({color, setShowSearchMenu}) {
    const [iconVisible, setIconVisible] = useState(true);
    const menu = useRef(null)
    const input = useRef(null)
    useClickOutside(menu, () => {
        setShowSearchMenu(false);
    })
    useEffect(() => {
        if (input.current) {
            input.current.focus();
        }
    }, [input])
    return <div className="header_left search_area scrollbar" ref={menu}>
        <div className="search_wrap">
            <div className="header_logo">
                <div className="circle hover1" onClick={() => setShowSearchMenu(false)}>
                    <Return color={color}/>
                </div>
            </div>
            <div className="search" onClick={() => {
                input.current.focus();
            }}>
                {iconVisible && <div><Search color={color}/></div>}
                <input type="text"
                       placeholder='Search Facebook'
                       onFocus={() => setIconVisible(false)}
                       rel={ref}
                       onBlur={() => setIconVisible(true)}
                />
            </div>
        </div>
        <div className="search_history_header">
            <span>Recent searches</span>
            <a>Edit</a>
        </div>
        <div className="search_history"></div>
        <div className="search_results scrollbar"></div>
    </div>

}

export default SearchMenu;