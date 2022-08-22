import axios from "axios";
import Cookies from 'js-cookie';

const axiosMain = axios.create({
    baseURL: `${process.env.REACT_APP_BACKEND_URL}/api/v1`,
    headers: {
        'Content-Type': 'application/json',
        'Access-Control-Allow-Origin': '*',
        'Access-Control-Allow-Methods': 'GET,PUT,POST,DELETE,PATCH,OPTIONS',
    },
})
axiosMain.interceptors.request.use(config => {
    const userCookie = Cookies.get("user")
    if(userCookie){
        const user = JSON.parse(userCookie)
        config.headers.Authorization = `Bearer ${user.token}`;
    }
    return config;
})

axiosMain.interceptors.response.use(response => {
    if (response.status === 200 || response.status === 201) {
        return response;
    }
    return Promise.reject(response);
})
export {axiosMain}

