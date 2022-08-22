import {Form, Formik} from "formik";
import {Link, Navigate, useNavigate} from "react-router-dom";
import LoginInput from "../../components/inputs/loginInput";
import * as Yup from "yup";
import {useState} from "react";
import {axiosMain} from "../../axios/axiosMain";
import PropagateLoader from "react-spinners/PropagateLoader";
import {extractErrorFromResponse} from "../../common";

export default function ChangePassword({
                                           password,
                                           setPassword,
                                           conf_password,
                                           setConf_password,
                                           error,
                                           token,
                                           setError,
                                           setLoading,
                                           loading,
                                           userInfo
                                       }) {
    const [success, setSuccess] = useState('')
    const navigate = useNavigate()
    const validatePassword = Yup.object({
        password: Yup.string()
            .required(
                "Enter a combination of at least six numbers,letters and punctuation marks(such as ! and &)."
            )
            .min(6, "Password must be at least 6 characters.")
            .max(36, "Password can't be more than 36 characters"),

        conf_password: Yup.string()
            .required("Confirm your password.")
            .oneOf([Yup.ref("password")], "Passwords must match."),
    });

    if (!token) {
        return <Navigate to='/'/>
    }

    async function handleSubmitChangePassword(values) {
        setLoading(true)
        setSuccess('')
        try {
            const {data} = await axiosMain.post('/users/changePassword', {
                password: values.password,
                conf_password: values.conf_password,
                token: token,
                email: userInfo.email
            })
            setError('')
            setSuccess(data.message)
            setTimeout(() => {
                navigate('/login')
            }, 2000)
        } catch (e) {
            setError(extractErrorFromResponse(e.response.data))
        } finally {
            setLoading(false)
        }
    }

    return (
        <div className="reset_form" style={{height: "310px"}}>
            <div className="reset_form_header">Change Password</div>
            <div className="reset_form_text">Pick a strong password</div>
            <Formik
                enableReinitialize
                initialValues={{
                    password,
                    conf_password,
                }}
                validationSchema={validatePassword}
                onSubmit={handleSubmitChangePassword}
            >
                {(formik) => (
                    <Form>
                        <LoginInput
                            type="password"
                            name="password"
                            onChange={(e) => setPassword(e.target.value)}
                            placeholder="New password"
                        />
                        <LoginInput
                            type="password"
                            name="conf_password"
                            onChange={(e) => setConf_password(e.target.value)}
                            placeholder="Confirm new password"
                            bottom
                        />
                        <div className='reset_notification'>
                            {error && <div className="error_text">{error}</div>}
                            {loading && <PropagateLoader color={'blue'}/>}
                            {success && <div className='success_text'>{success}</div>}
                        </div>

                        <div className="reset_form_btns">
                            <Link to="/login" className="gray_btn">
                                Cancel
                            </Link>
                            <button type="submit" className="blue_btn">
                                Continue
                            </button>
                        </div>
                    </Form>
                )}
            </Formik>
        </div>
    );
}
