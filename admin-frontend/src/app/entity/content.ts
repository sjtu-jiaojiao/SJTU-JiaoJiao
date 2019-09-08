export class Media{
    fileID: string;
    type: number;
}

export class ContentResponse
{
    status: string;
    files: Media[];
    tags: string[];
}