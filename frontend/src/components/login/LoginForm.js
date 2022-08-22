import {Form, Formik} from "formik";
import LoginInput from "../inputs/loginInput";
import {Link, useNavigate} from "react-router-dom";
import {useState} from "react";
import * as Yup from "yup";
import {useDispatch} from "react-redux";
import Cookies from "js-cookie";
import {axiosMain} from "../../axios/axiosMain";
import {DotLoader} from "react-spinners";

const loginInfos = {
    email: "",
    password: "",
};

const loginValidation = Yup.object({
    email: Yup.string()
        .required("Email address is required.")
        .email("Must be a valid email.")
        .max(100),
    password: Yup.string().required("Password is required"),
});


function LoginForm({setShowRegister}) {
    const [login, setLogin] = useState(loginInfos);
    const {email, password} = login;
    const navigate = useNavigate()
    const dispatch = useDispatch();
    const [error, setError] = useState('');
    const [success, setSuccess] = useState('')
    const [loading, setLoading] = useState(false);

    const handleLoginChange = (e) => {
        const {name, value} = e.target;
        setLogin({...login, [name]: value});
    };

    async function loginSubmit() {
        setLoading(true)
        try {
            let {data} = await axiosMain.post(`/users/login`, {
                email,
                password
            });
            setError('')
            setSuccess(data.message)
            dispatch({type: 'LOGIN', payload: data.data});
            Cookies.set('user', JSON.stringify(data.data));
            navigate('/')
        } catch (error) {
            console.log(error)
            setSuccess('')
            const data = error.response.data;
            const errors = data.errors;
            const msg = data.message;

            if (msg) {
                setError(msg)
            }
            if (errors.length > 0) {
                const errorsString = errors.map(error => JSON.stringify(error)).join(', ');
                setError(errorsString)
            }
        } finally {
            setLoading(false)
        }
    }

    return <div className="login_wrap">
        <div className="login_1">
            <img src="../../icons/facebook.svg" alt=""/>
            <span>
              Facebook helps you connect and share with the people in your life.
            </span>
        </div>
        <div className="login_2">
            <div className="login_2_wrap">
                <Formik
                    enableReinitialize
                    initialValues={{
                        email,
                        password,
                    }}
                    validationSchema={loginValidation}
                    onSubmit={loginSubmit}
                >
                    {(formik) => (
                        <Form>
                            <LoginInput
                                type="text"
                                name="email"
                                placeholder="Email address or phone number"
                                onChange={handleLoginChange}
                            />
                            <LoginInput
                                type="password"
                                name="password"
                                placeholder="Password"
                                onChange={handleLoginChange}
                                bottom
                            />
                            <button type="submit" className="blue_btn">
                                Log In
                            </button>
                        </Form>
                    )}
                </Formik>
                <Link to="/reset" className="forgot_password">
                    Forgotten password?
                </Link>
                <div className="sign_splitter"></div>
                <button onClick={() => setShowRegister(true)} className="blue_btn open_signup">Create Account</button>
                <DotLoader loading={loading} color='#1875f2' size={30}/>
                {error && <div className='error_text'>{error}</div>}
                {success && <div className='success_text'>{success}</div>}
            </div>
            <Link to="/" className="sign_extra">
                <b>Create a Page</b> for a celebrity, brand or business.
            </Link>
        </div>
    </div>;
}

export default LoginForm;