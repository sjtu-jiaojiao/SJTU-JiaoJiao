export const isUserNameValid = (userName) => {
    //return /^[a-zA-Z]{1}([a-zA-Z0-9]|[._]){0,31}$/.test(this.state.searchText);
    return /^([A-Za-z0-9\u4e00-\u9fa5]|[._]){1,32}$/gi.test(userName);
};

export const isTelephoneValid = (telephone) => {
    let length = telephone.length;
    return length === 11 && /^(((13[0-9]{1})|(15[0-9]{1})|(18[0-9]{1})|(14[0-9]{1})|)+\d{8})$/.test(telephone);
};

export const isValidTimeValid = (validTime) => {
    return /^[1-9]\d*$/.test(validTime);
};

export const isPriceValid = (Price) => {
    return /^(([1-9][0-9]*)|(([0]\.\d{1,2}|[1-9][0-9]*\.\d{1,2})))$/.test(Price);
};
