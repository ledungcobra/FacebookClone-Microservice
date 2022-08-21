import {useSelector} from "react-redux";
import Login from "../pages/login";
import {Outlet} from "react-router-dom";


export default function LoggedInRoutes({element}) {
    const user = useSelector(state => state.user);
    return user ? <Outlet/> : <Login/>
}