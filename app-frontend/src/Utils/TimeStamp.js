export const TimeStamptoDate = (time) => {
    let datetime = new Date();
    datetime.setTime(time * 1000);
    let year = datetime.getFullYear();
    let month = datetime.getMonth() + 1;
    let date = datetime.getDate();
    let hour = datetime.getHours();
    let minute = datetime.getMinutes();
    let second = datetime.getSeconds();
    return year + "-" + month + "-" + date + " " + hour + ":" + minute + ":" + second;
};

export const DatetoTimeStamp = (date) => {
    let d = new Date(date);
    return Math.round(d.getTime());
};

export const TimeStampNow = () => {
    return new Date().getTime();
};
