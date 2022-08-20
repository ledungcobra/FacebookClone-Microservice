import {Form, Formik} from "formik";
import {useState} from "react";
import RegisterInput from "../inputs/registerInput";
import * as Yup from 'yup'
import DateOfBirthSelect from "./DateOfBirthSelect";
import GenderSelect from "./GenderSelect";
import {DotLoader} from "react-spinners";
import {useDispatch} from "react-redux";
import Cookies from "js-cookie";
import {useNavigate} from 'react-router-dom'
import PropTypes from 'prop-types';
import {axiosMain} from "../../axios/axiosMain";

const now = new Date();
const userInfos = {
    first_name: "",
    last_name: "",
    email: "",
    password: "",
    bYear: now.getFullYear(),
    bMonth: now.getMonth() + 1,
    bDay: now.getDate(),
    gender: "",
};
let nameSchema = Yup.string().min(2, 'First name must between 2 and 16 characters.')
    .max(16, 'First name must between 2 and 16 characters.')
    .matches(/^([aA-zZ]\s?)+$/, "Number or special character is not allowed.");

const registerValidation = () => Yup.object({
    first_name: nameSchema.required('What\'s your first name?'),
    last_name: nameSchema.required('What\'s your last name?'),
    email: Yup.string()
        .required('What\'s your email address?')
        .email('Enter a valid email'),
    password: Yup.string()
        .required('What\'s your password?')
        .min(8, 'Password must between 8 and 16 characters.')
        .max(16, 'Password must between 8 and 16 characters.')

});


export default function RegisterForm({hideForm}) {
    const [user, setUser] = useState(userInfos);
    const [error, setError] = useState('');
    const [success, setSuccess] = useState('')
    const [loading, setLoading] = useState(false);
    const dispatch = useDispatch();
    const navigate = useNavigate();

    const handleRegisterChange = (e) => {
        const {name, value} = e.target;
        setUser({...user, [name]: value});
    };

    async function registerSubmit(user) {
        setLoading(true)
        try {
            let {data} = await axiosMain.post(`/users/register`, {
                ...user,
                birth_day: +user.bDay,
                birth_month: +user.bMonth,
                birth_year: +user.bYear,
            });
            setError('')
            setSuccess(data.message)
            setTimeout(() => {
                dispatch({type: 'LOGIN', payload: data.data});
                Cookies.set('user', JSON.stringify(data.data));
                navigate('/')
            }, 2000)
        } catch (error) {
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

    return (<div className="blur">
        <div className="register">
            <div className="register_header">
                <i onClick={hideForm} className="exit_icon"></i>
                <span>Sign Up</span>
                <span>it's quick and easy</span>
            </div>
            <Formik enableReinitialize
                    initialValues={user}
                    validationSchema={registerValidation}
                    onSubmit={() => {
                        registerSubmit(user);
                    }}
            >
                {(formik) => (<Form className="register_form" onChange={handleRegisterChange}>
                    <div className="reg_line">
                        <RegisterInput
                            type="text"
                            placeholder="First name"
                            name="first_name"
                            onChange={handleRegisterChange}
                        />
                        <RegisterInput
                            type="text"
                            placeholder="Surname"
                            name="last_name"
                            onChange={handleRegisterChange}
                        />
                    </div>
                    <div className="reg_line">
                        <RegisterInput
                            type="text"
                            placeholder="Mobile number or email address"
                            name="email"
                            onChange={handleRegisterChange}
                        />
                    </div>
                    <div className="reg_line">
                        <RegisterInput
                            type="password"
                            placeholder="New password"
                            name="password"
                            onChange={handleRegisterChange}
                        />
                    </div>
                    <div className="reg_col">
                        <div className="reg_line_header">
                            Date of birth <i className="info_icon"></i>
                        </div>
                        <DateOfBirthSelect
                            bDay={user.bDay}
                            bYear={user.bYear}
                            bMonth={user.bMonth}
                            handleRegisterChange={handleRegisterChange}/>
                    </div>
                    <div className="reg_col">
                        <div className="reg_line_header">
                            Gender <i className="info_icon"></i>
                        </div>
                        <GenderSelect handleRegisterChange={handleRegisterChange}/>
                    </div>
                    <div className="reg_infos">
                        By clicking Sign Up, you agree to our{" "}
                        <span>Terms, Data Policy &nbsp;</span>
                        and <span>Cookie Policy.</span> You may receive SMS
                        notifications from us and can opt out at any time.
                    </div>
                    <div className="reg_btn_wrapper">
                        <button className="blue_btn open_signup">Sign Up</button>
                    </div>
                    <DotLoader loading={loading} color='#1875f2' size={30}/>
                    {error && <div className='error_text'>{error}</div>}
                    {success && <div className='success_text'>{success}</div>}

                </Form>)}
            </Formik>
        </div>
    </div>);
}

RegisterForm.PropsTypes = {
    hideForm: PropTypes.func.isRequired,
}