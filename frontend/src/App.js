import {BrowserRouter, Route, Routes} from "react-router-dom";
import Login from "./pages/login";
import Profile from "./pages/profile";
import Home from "./pages/home";
import {Provider} from "react-redux";
import {createStore} from "redux";
import {composeWithDevTools} from "redux-devtools-extension";
import rootReducer from "./reducers";

const store = createStore(rootReducer, composeWithDevTools());

function App() {
    return (
        <Provider store={store}>
            <BrowserRouter>
                <Routes>
                    <Route path="/login" element={<Login/>} exact/>
                    <Route path="/profile" element={<Profile/>} exact/>
                    <Route path="/" element={<Home/>} exact/>
                </Routes>
            </BrowserRouter>
        </Provider>
    );
}

export default App;
