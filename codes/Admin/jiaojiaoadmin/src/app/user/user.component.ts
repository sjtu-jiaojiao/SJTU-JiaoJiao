import { Component, OnInit } from '@angular/core';
import { User } from '../entity/user';
import { UserService } from '../user.service';

@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.css']
})
export class UserComponent implements OnInit {
  users: User[];
  threshold: number;
  current : number;
  curusers: User[];
  size :number;
  count : number;
  constructor(private userService: UserService) { }

  ngOnInit() {
    this.getusers();
  }

  getusers(): void {
    this.userService.getUsers()
    .subscribe(users => {
      this.users = users;
      this.current=1;
      this.size=4;
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
    this.userService.deleteUser(user).subscribe(_ =>{
      this.count = this.users.length; 
      this.switchPage(this.current, this.size);
    });
  }
  add(id: string): void {
    id = id.trim();
    if (!id) { return; }
    this.userService.addUser({ id } as User)
      .subscribe(user => {
        this.users.push(user);
      });
  }

}
