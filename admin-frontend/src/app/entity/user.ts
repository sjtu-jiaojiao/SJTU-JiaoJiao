export class User
{
   userId: number;
   userName: string;
   avatarId: string;
   telephone: string;
   studentId: string;
   studentName: string;
   status: number;
   role: number;
}

export class UserResponse
{
   user: User[];
}