import {axiosMain} from "../axios/axiosMain";

export const createPost = async (type,
                                 background,
                                 text,
                                 images) => {
    try {
        console.log(images)
        const {data} = await axiosMain.post('/posts', {
            type, background, text, images,
        })
        return {success: true, data: data.data, message: data.message}
    } catch (e) {
        console.log(e)
        return {message: e.response.data.message || e.message, success: false}
    }
}
