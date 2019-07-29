export class sellInfo{
    public sellInfoID : number;
    public status : number;
    public releaseTime: string;
    public validTime: string;
    public goodName: string;
    public price: number;
    public description: string;
    public contentID: string;
    public userID: number;
    public tag: string [];
}

export class buyInfo{
    public buyInfoID : number;
    public status : number;
    public releaseTime: string;
    public validTime: string;
    public goodName: string;
    public price: number;
    public description: string;
    public contentID: string;
    public userID: number;
    public tag: string [];
}

export class InfoResponse{
    public sellInfo: sellInfo[];
    public buyInfo: buyInfo[];
}