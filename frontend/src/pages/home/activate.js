import React, {useEffect, useState} from 'react';
import Header from "../../components/header";
import LeftHome from "../../components/home/left";
import {useDispatch, useSelector} from "react-redux";
import RightHome from "../../components/home/right";
import Stories from "../../components/home/stories";
import './style.css'
import CreatePost from "../../components/createPost";
import ActivateForm from "./ActivateForm";
import useQuery from "../../hooks/useQuery";
import {axiosMain} from "../../axios/axiosMain";
import Cookies from 'js-cookie'
import {useNavigate} from "react-router-dom";
import {VERIFIED} from "../../common/constants";

function Activate() {
    const user = useSelector(state => state.user);

    const [success, setSuccess] = useState('');
    const [error, setError] = useState('')
    const [loading, setLoading] = useState(false);
    const dispatch = useDispatch();
    const navigate = useNavigate();

    const query = useQuery()
    const token = query.get('token')
    const email = query.get('email')
    useEffect(() => {
        activateAccount();
    }, [])

    async function activateAccount() {
        setLoading(true)
        try {
            const {data} = await axiosMain.post('/users/activate', {
                email,
                token
            })
            setSuccess(data.message)
            Cookies.set('user', JSON.stringify({...user, verified: true}))
            dispatch({type: VERIFIED, verified: true})
        } catch (e) {
            setError(e.response.data.message)
        } finally {
            setTimeout(() => {
                navigate('/')
                setLoading(false)
            }, 2000)
        }
    }

    return (
        <div className='home'>
            {success && (
                <ActivateForm
                    type="success"
                    header="Account verification succeded."
                    text={success}
                    loading={loading}
                />
            )}
            {error && (
                <ActivateForm
                    type="error"
                    header="Account verification failed."
                    text={error}
                    loading={loading}
                />
            )}
            <Header user={user}/>
            <LeftHome user={user}/>
            <div className="home_middle">
                <Stories/>
                <CreatePost user={user}/>
            </div>
            <RightHome user={user}/>
        </div>
    );
}

export default Activate;