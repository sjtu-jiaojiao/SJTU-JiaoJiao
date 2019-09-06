let Config = {
    fetchPrefix: 'http://202.120.40.8:30711/v1/',
    JaccountToken: {},
    userInfo: {
        userID: -1,
        userName: '',
        avatarID: 'to_be_changed',
        telephone: '',
        studentID: '',
        studentName: '',
    },
    isReleaseRender: false,
    isContactRender: false,
};

export const isLogin = () => {
    return Config.userInfo.userID !== -1;
};

export const isReleaseRendered = () => {
    return Config.isReleaseRender;
};

export const isContactRendered = () => {
    return Config.isContactRender;
};

export default Config;
