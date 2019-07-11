import { Component } from '@angular/core';
import { AuthService } from './auth.service';
import { NzNotificationService } from 'ng-zorro-antd';
import { Router } from '@angular/router';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
})
export class AppComponent {
  title = 'JOJO-Admin';
  isCollapsed = false;
  constructor( private router: Router, private authService: AuthService, private notification: NzNotificationService){}
  logout(): boolean{
    if(this.authService.isLoggedIn){
    this.authService.logout();
    this.notification.create('success', '注销成功', '退出登录');
    // Redirect the user
    this.router.navigateByUrl('/login');
    return true;
    }
    else 
    return false;
  }
}
