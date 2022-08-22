import {Form, Formik} from "formik";
import {Link} from "react-router-dom";
import LoginInput from "../../components/inputs/loginInput";
import * as Yup from "yup";
import {axiosMain} from "../../axios/axiosMain";
import {useState} from "react";
import PropagateLoader from "react-spinners/PropagateLoader";

export default function CodeVerification({code, setCode, error, setToken, userInfo, setError, loading, setLoading,setVisible}) {
    const [success, setSuccess] = useState('')
    const validateCode = Yup.object({
        code: Yup.string()
            .required("Code is required")
            .min("5", "Code must be 5 characters.")
            .max("5", "Code must be 5 characters."),
    });

    async function handleCodeVerification() {
        setLoading(true)
        setSuccess('')
        try {
            const {data} = await axiosMain.post('/users/verifyCode', {
                email: userInfo.email,
                code: code
            })
            setToken(data.data.token)
            setError('')
            setSuccess(data.message)
            setVisible(3)
        } catch (e) {
            const data = e.response.data
            setError(data.message)
        } finally {
            setLoading(false)
        }
    }

    return (
        <div className="reset_form">
            <div className="reset_form_header">Code verification</div>
            <div className="reset_form_text">
                Please enter code that been sent to your email.
            </div>
            <Formik
                enableReinitialize
                initialValues={{
                    code,
                }}
                validationSchema={validateCode}
            >
                {(formik) => (
                    <Form>
                        <LoginInput
                            type="text"
                            name="code"
                            onChange={(e) => setCode(e.target.value)}
                            placeholder="Code"
                        />
                        {error && <div className="error_text">{error}</div>}
                        {loading && <PropagateLoader/>}
                        {success && <div className='success_text'>{success}</div>}
                        <div className="reset_form_btns">
                            <Link to="/login" className="gray_btn">
                                Cancel
                            </Link>
                            <button onClick={handleCodeVerification} type="submit" className="blue_btn">
                                Continue
                            </button>
                        </div>
                    </Form>
                )}
            </Formik>
        </div>
    );
}
