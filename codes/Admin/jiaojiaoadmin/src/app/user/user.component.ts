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
  constructor(private userService: UserService) { }

  ngOnInit() {
    this.getusers();
  }

  getusers(): void {
    this.userService.getUsers()
    .subscribe(users => this.users = users);
  }

  forbid(): void {
    this.users.filter(m => m.score < this.threshold).map( m => {m.forbid = true; return m;}).forEach(element => 
      this.userService.updateUser(element).subscribe());
    this.users.filter(m => m.score >= this.threshold).map( m => {m.forbid = false; return m;}).forEach(element => 
        this.userService.updateUser(element).subscribe());
  }

  delete(user: User): void {
    this.users = this.users.filter(h => h !== user);
    this.userService.deleteUser(user).subscribe();
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
