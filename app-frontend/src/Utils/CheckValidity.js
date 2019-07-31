export const isUserNameValid = (userName) => {
    //return /^[a-zA-Z]{1}([a-zA-Z0-9]|[._]){0,31}$/.test(this.state.searchText);
    return /^([A-Za-z0-9\u4e00-\u9fa5]|[._]){1,32}$/gi.test(userName);
};
