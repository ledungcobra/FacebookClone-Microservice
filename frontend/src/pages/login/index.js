import "./style.css";
import LoginForm from "../../components/login/LoginForm";
import Footer from "../../components/login/Footer";
import RegisterForm from "../../components/login/RegisterForm";
import {useState} from "react";


export default function Login() {

    const [showRegister, setShowRegister] = useState(false);

    return (
        <div className="login">
            <div className="login_wrapper">
                <LoginForm setShowRegister={setShowRegister}/>
                {showRegister && <RegisterForm hideForm={()=>setShowRegister(false)}/>}
                <Footer/>
            </div>
        </div>
    );
}
