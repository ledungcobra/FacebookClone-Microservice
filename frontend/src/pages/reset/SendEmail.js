import {Link} from "react-router-dom";
import {axiosMain} from "../../axios/axiosMain";
import PropagateLoader from "react-spinners/PropagateLoader";
import {useState} from "react";

export default function SendEmail({userInfo, error, setError, setVisible, setLoading, loading}) {

    const [success, setSuccess] = useState('')
    if (!userInfo) {
        return null
    }

    async function handleResetAccount() {
        setLoading(true)
        try {
            const {data} = await axiosMain.post('/users/resetPassword', {
                email: userInfo.email
            });
            setError('')
            setSuccess(data.message)
            setTimeout(() => {
                setVisible(2)
            }, 2000)
        } catch (e) {
            setError(e.response.data.message)
        } finally {
            setLoading(false)
        }
    }

    return (<div className="reset_form dynamic_height">
        <div className="reset_form_header">Reset Your Password</div>
        <div className="reset_grid">
            <div className="reset_left">
                <div className="reset_form_text">
                    How do you want to receive the code to reset your password?
                </div>
                <label htmlFor="email" className="hover1">
                    <input type="radio" name="" id="email" checked readOnly/>
                    <div className="label_col">
                        <span>Send code via email</span>
                        <span>{userInfo.email}</span>
                    </div>
                </label>
            </div>
            <div className="reset_right">
                <img src={userInfo.picture} alt=""/>
                <span>{userInfo.email}</span>
                <span>Facebook user</span>
            </div>
        </div>
        <div className='reset_notification'>
            {error && <div className='error_text'>
                {error}
            </div>}
            {success && <div className='success_text'>
                {success}
            </div>}
            {loading && <PropagateLoader color={'blue'}/>}
        </div>
        <div className="reset_form_btns">
            <Link to="/login" className="gray_btn">
                Not You ?
            </Link>
            <button onClick={handleResetAccount} type="submit" className="blue_btn">
                Continue
            </button>
        </div>
    </div>);
}
