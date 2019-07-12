import { Component, OnInit } from '@angular/core';
import { User } from '../entity/user';
import { UserService } from '../user.service';
import { HTTP_INTERCEPTORS } from '@angular/common/http';
import { JWTInterceptor } from '@delon/auth';

@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.css'],    
  //providers: [{ provide: HTTP_INTERCEPTORS, useClass: JWTInterceptor, multi: true}]

})
export class UserComponent implements OnInit {
  users: User[];
  threshold: number;
  current : number =1 ;
  curusers: User[] ; 
  size :number = 4;
  count : number;
  searchName: string;
  constructor(private userService: UserService) { }

  ngOnInit() {
    this.getusers();
  }
  searchByName(): void {
    this.userService.searchUsers(this.searchName)
    .subscribe(users => {
      this.users = users;
      this.count = this.users.length; 
      this.switchPage(this.current, this.size);
    });
  }
  getusers(): void {
    this.userService.getUsers()
    .subscribe(users => {
      this.users = users;
      this.count = this.users.length; 
      this.switchPage(this.current, this.size);
    });
  }
  switchPage(page, size) {
    if(page* size< this.count)
    this.curusers = this.users.slice( (page-1)*size, page* size)
    else 
    this.curusers = this.users.slice( (page-1)*size );
  }
  pageChange(page){
    this.switchPage(page,this.size);
  }
  sizeChange(size){
    this.switchPage(this.current,size);
  }
  forbid(): void {
    this.users.filter(m => m.score < this.threshold).map( m => {m.forbid = true; return m;}).forEach(element => 
      this.userService.updateUser(element).subscribe());
    this.users.filter(m => m.score >= this.threshold).map( m => {m.forbid = false; return m;}).forEach(element => 
        this.userService.updateUser(element).subscribe());
  }

  delete(user: User): void {
    this.users = this.users.filter(h => h !== user);
    this.userService.deleteUser(user.id).subscribe(_ =>{
      this.count = this.users.length; 
      this.switchPage(this.current, this.size);
    });
  }
}
