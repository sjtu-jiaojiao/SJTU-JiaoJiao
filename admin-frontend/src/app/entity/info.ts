export class Info{
    public sellInfoId : number;
    public status : number;
    public realeaseTime: string;
    public validTime: string;
    public goodName: string;
    public price: number;
    public description: string;
    public contentId: string;
    public userId: number;
    public tag: string [];
}

export class InfoResponse{
    public sellInfo: Info[];
}