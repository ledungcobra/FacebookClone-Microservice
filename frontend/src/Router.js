import {BrowserRouter, Route, Routes} from "react-router-dom";
import Profile from "./pages/profile";
import Home from "./pages/home";
import {Provider} from "react-redux";
import {createStore} from "redux";
import {composeWithDevTools} from "redux-devtools-extension";
import rootReducer from "./reducers";
import LoggedInRoutes from "./routes/LoggedInRoutes";
import Activate from "./pages/home/activate";
import NotLoggedInRoutes from "./routes/NotLoggedInRoutes";
import Login from "./pages/login";
import Reset from "./pages/reset";

const store = createStore(rootReducer, composeWithDevTools());

function Router() {
    return (
        <Provider store={store}>
            <BrowserRouter>
                <Routes>
                    <Route element={<LoggedInRoutes/>}>
                        <Route path="/" element={<Home/>} exact/>
                        <Route path="/profile" element={<Profile/>} exact/>
                        <Route path="/activate" element={<Activate/>} exact/>
                    </Route>
                    <Route element={<NotLoggedInRoutes/>}>
                        <Route path="/login" element={<Login/>} exact/>
                    </Route>
                    <Route path='/reset' element={<Reset />}/>
                </Routes>
            </BrowserRouter>
        </Provider>
    );
}

export default Router;
