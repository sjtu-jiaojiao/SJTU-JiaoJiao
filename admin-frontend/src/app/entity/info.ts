export class sellInfo{
    public sellInfoID : number;
    public status : number;
    public releaseTime: number;
    public validTime: number;
    public goodName: string;
    public price: number;
    public description: string;
    public contentID: string;
    public userID: number;
    public tags: string [];
}

export class buyInfo{
    public buyInfoID : number;
    public status : number;
    public releaseTime: number;
    public validTime: number;
    public goodName: string;
    public price: number;
    public description: string;
    public contentID: string;
    public userID: number;
    public tags: string [];
}

export class InfoResponse{
    public sellInfo: sellInfo[];
    public buyInfo: buyInfo[];
}