import {Form, Formik} from "formik";
import {Link} from "react-router-dom";
import LoginInput from "../../components/inputs/loginInput";
import * as Yup from "yup";
import {axiosMain} from "../../axios/axiosMain";

export default function SearchAccount({email, setEmail, error, setError, setLoading, setUserInfo, setVisible}) {

    const validateEmail = Yup.object({
        email: Yup.string()
            .required("Email address ir required.")
            .email("Must be a valid email address.")
            .max(50, "Email address can't be more than 50 characters."),
    });


    async function handleSearch(formData) {
        setLoading(true)
        try {
            const {email} = formData;
            const {data} = await axiosMain.get('/users?email=' + email)
            setUserInfo(data.data)
            setVisible(1)
            setError('')
        } catch (e) {
            setError(e.response.data.message)
        } finally {
            setLoading(false)

        }
    }

    return (
        <div className="reset_form">
            <div className="reset_form_header">Find Your Account</div>
            <div className="reset_form_text">
                Please enter your email address or mobile number to search for your
                account.
            </div>
            <Formik
                enableReinitialize
                initialValues={{
                    email,
                }}
                onSubmit={handleSearch}
                validationSchema={validateEmail}
            >
                {(formik) => (
                    <Form>
                        <LoginInput
                            type="text"
                            name="email"
                            onChange={(e) => setEmail(e.target.value)}
                            placeholder="Email address or phone number"
                        />
                        {error && <div className="error_text">{error}</div>}
                        <div className="reset_form_btns">
                            <Link to="/login" className="gray_btn">
                                Cancel
                            </Link>
                            <button type="submit" className="blue_btn">
                                Search
                            </button>
                        </div>
                    </Form>
                )}
            </Formik>
        </div>
    );
}
