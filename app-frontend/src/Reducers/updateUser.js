import {
    UPDATE_TELEPHONE,
    UPDATE_USERNAME,
    UPDATE_USERINFO,
} from "../Actions/types";

let initState = {
    userId: -1,
    userName: '',
    studentId: '',
    studentName: '',
    telephone: ''
};

const updateUser = (state=initState, action) => {
    switch (action.type) {
        case UPDATE_USERINFO:
            return {
                userId: action.userId,
                userName: action.userName,
                studentId: action.studentId,
                studentName: action.studentName,
                telephone: action.telephone
            };
        case UPDATE_USERNAME:
            return {
                ...state,
                userName: action.userName
            };
        case UPDATE_TELEPHONE:
            return {
                ...state,
                telephone: action.telephone
            };
        default:
            return state;
    }
};

export default updateUser;
