import { Component, OnInit, Inject } from '@angular/core';
import { SocialService, JWTGuard, JWTTokenModel } from '@delon/auth';
import { AuthService } from '../auth.service';
import { Router } from '@angular/router';
import { NzNotificationService } from 'ng-zorro-antd';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
  providers: [ SocialService ]

})
export class LoginComponent implements OnInit {

  t: any;
  constructor( private socialService: SocialService , private notification: NzNotificationService
    ,private authService: AuthService) {
  }

  ngOnInit(): void {
  }

  
 login() {
    this.socialService.login('/api/auth','/callback',{type: 'window'} ).subscribe(res => {
//      this.authService.login(res);
    })
  }

  pretend() {
    this.authService.login({token: this.t});
    this.notification.create('success', '导入token成功', 'Token已更新');
  }

  }