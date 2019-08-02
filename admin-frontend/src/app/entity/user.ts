export class User
{
   userID: number;
   userName: string;
   avatarID: string;
   telephone: string;
   studentID: string;
   studentName: string;
   status: number;
   role: number;
}

export class UserResponse
{
   user: User[];
}