export class User
{
   userId: number;
   userName: string;
   avatarId: string;
   telephone: string;
   studentId: string;
   studentName: string;
   status: number;
}

export class UserResponse
{
   user: User[];
}