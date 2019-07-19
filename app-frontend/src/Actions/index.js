import {
    UPDATE_TELEPHONE,
    UPDATE_USERNAME,
    UPDATE_USERINFO,
} from "./types";
import Config from "../Config";

export const updateUserInfo = (responseJson) => ({
    type: UPDATE_USERINFO,
    userId: responseJson.userId,
    userName: responseJson.userName,
    studentId: responseJson.studentId,
    studentName: responseJson.studentName,
    telephone: responseJson.telephone
});

export const updateUserName = (userName) => ({
    type: UPDATE_USERNAME,
    userName: userName
});

export const updateUserTelephone = (telephone) => ({
    type: UPDATE_TELEPHONE,
    telephone: telephone
});
