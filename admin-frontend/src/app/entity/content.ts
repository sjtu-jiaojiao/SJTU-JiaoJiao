export class File{
    fileId: string;
    type: number;
}

export class ContentResponse
{
    status: string;
    file: File[];
}