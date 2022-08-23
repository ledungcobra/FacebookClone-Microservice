import {axiosMain} from "../axios/axiosMain";

export const uploadImage = async (imagesBlob) => {
    try {
        const form = new FormData()
        imagesBlob.map(blob => form.append('file', blob))
        const {data} = await axiosMain.post('/posts/uploadImages', form, {
            headers: {
                'Content-Type': 'multipart/form-data'
            }
        })
        return {
            success: true,
            images: data.data.results.map(r => r.url),
            message: 'Upload images successfully'
        }
    } catch (e) {
        return {success: false, message: e.response?.data?.message || e?.message}
    }
}
