import axios from "axios";

const axiosMain = axios.create({
    baseURL: `${process.env.REACT_APP_BACKEND_URL}/api/v1`,
    headers: {
        'Content-Type': 'application/json',
        'Access-Control-Allow-Origin': '*',
        'Access-Control-Allow-Methods': 'GET,PUT,POST,DELETE,PATCH,OPTIONS',
    },
})
axiosMain.interceptors.request.use(config => {
    const token = localStorage.getItem('token');
    if (token) {
        config.headers.Authorization = `Bearer ${token}`;
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

