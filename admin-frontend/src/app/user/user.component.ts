import { Component, OnInit } from '@angular/core';
import { User } from '../entity/user';
import { UserService } from '../user.service';
import { HTTP_INTERCEPTORS } from '@angular/common/http';
import { JWTInterceptor } from '@delon/auth';

@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.css']
})
export class UserComponent implements OnInit {
  users: User[];
  threshold: number;
  current : number = 1;
  count: number;
  size: number = 4;
  searchName: string;
  constructor(private userService: UserService) { }

  ngOnInit() {
    this.getusers();
  }

  searchByName(): void {    
    if (!this.searchName || !this.searchName.trim()) {
    this.getusers();
    // if not search term, return all user array.
    return;
  }
    this.userService.searchUsers(this.searchName,this.size, this.current*this.size-this.size)
    .subscribe(users => {
      if(!users) return;
      this.users = users.user;
      this.checkcount();
    });
  }

  checkcount(){
    if(this.users && this.users.length === this.size)
    this.count = (this.current + 1) * this.size;
    else 
    this.count = this.current * this.size;
  }
  getusers(): void {
    this.userService.getPageUsers(this.size, this.current*this.size-this.size)
    .subscribe(users => {
      if(!users) return;
      this.users = users.user;
      this.checkcount();
    });
  }

  onChange(){
    this.searchByName();
  }
  /*forbid(): void {
    this.users.filter(m => m.score < this.threshold).map( m => {m.forbid = true; return m;}).forEach(element => 
      this.userService.updateUser(element).subscribe());
    this.users.filter(m => m.score >= this.threshold).map( m => {m.forbid = false; return m;}).forEach(element => 
        this.userService.updateUser(element).subscribe());
  }*/

}
