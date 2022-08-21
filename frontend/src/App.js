import {BrowserRouter, Route, Routes} from "react-router-dom";
import Login from "./pages/login";
import Profile from "./pages/profile";
import Home from "./pages/home";
import {Provider} from "react-redux";
import {createStore} from "redux";
import {composeWithDevTools} from "redux-devtools-extension";
import rootReducer from "./reducers";
import LoggedInRoutes from "./routes/LoggedInRoutes";
import NotLoggedInRoutes from "./routes/NotLoggedInRoutes";

const store = createStore(rootReducer, composeWithDevTools());

function App() {
    return (
        <Provider store={store}>
            <BrowserRouter>
                <Routes>
                    <Route element={<LoggedInRoutes/>}>
                        <Route path="/" element={<Home/>} exact/>
                        <Route path="/profile" element={<Profile/>} exact/>
                    </Route>
                    <Route element={<NotLoggedInRoutes/>}>
                        <Route path="/login" element={<Login/>} exact/>
                    </Route>
                </Routes>
            </BrowserRouter>
        </Provider>
    );
}

export default App;
