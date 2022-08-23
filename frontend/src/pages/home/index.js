import React, {useState} from 'react';
import Header from "../../components/header";
import LeftHome from "../../components/home/left";
import {useSelector} from "react-redux";
import RightHome from "../../components/home/right";
import Stories from "../../components/home/stories";
import './style.css'
import CreatePost from "../../components/createPost";
import SendVerification from "../../components/sendVerification";
import CreatePostPopup from "../../components/createPostPopup";

function Home() {
    const user = useSelector(state => state.user);
    const [visible, setVisible] = useState(false)
    if (!user) return null

    return (
        <div className='home'>
            <Header user={user}/>
            <LeftHome user={user}/>
            <div className="home_middle">
                <Stories/>
                {!user.verified && <SendVerification user={user}/>}
                <CreatePost  user={user} setCreatePostVisible={setVisible}/>
            </div>
            <RightHome user={user}/>
            {visible && <CreatePostPopup setVisible={setVisible} user={user}/>}
        </div>
    );
}

export default Home;