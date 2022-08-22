import Cookies from "js-cookie";

const userReducer = function (state = Cookies.get('user') ? JSON.parse(Cookies.get('user')) : null, action) {
    switch (action.type) {
        case 'LOGIN':
            return action.payload;
        case 'VERIFIED':
            return {...state, verified: action.payload};
        case 'LOGOUT':
            Cookies.remove('user')
            return null
        default:
            return state;
    }
}

export default userReducer;
