import {TimeStamptoDate} from "./TimeStamp";

export const parseStatus = (status) => {
    switch (status) {
        case 1:
            return ('出售中');
        case 2:
            return ('已预约');
        case 3:
            return ('已出售');
        case 4:
            return ('已过期 (不再出售)');
    }
};

export const parseTimeStamp = (TimeStamp) => {
    let date = TimeStamptoDate(TimeStamp);
    return (date);
};

