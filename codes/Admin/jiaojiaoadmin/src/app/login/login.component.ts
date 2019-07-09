import { Component, OnInit, Inject } from '@angular/core';
import { Location } from '@angular/common';
import {
  FormBuilder,
  FormGroup,
  Validators
} from '@angular/forms';

import { NzNotificationService } from 'ng-zorro-antd';
import { Router } from '@angular/router';
import { AuthService } from '../auth.service';
import { AdminService } from '../admin.service';
//import { ITokenService, DA_SERVICE_TOKEN } from '@delon/auth';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  validateForm: FormGroup;
  submitForm(): void {
    for (const i in this.validateForm.controls) {
      if (this.validateForm.controls.hasOwnProperty(i)) {
        this.validateForm.controls[ i ].markAsDirty();
        this.validateForm.controls[ i ].updateValueAndValidity();
      }
    }
  }

  constructor(
    private authService: AuthService, public router: Router, 
    private adminService: AdminService,
    private fb: FormBuilder, private notification: NzNotificationService) {
  }

  ngOnInit(): void {
    this.validateForm = this.fb.group({
      userName: [ null, [ Validators.required ] ],
      password: [ null, [ Validators.required ] ]
    });
  }
  login(){
  this.adminService.login(this.validateForm.controls.userName.value , this.validateForm.controls.password.value).subscribe
  (
     e => {
  if(e && this.adminService.login(this.validateForm.controls.userName.value , this.validateForm.controls.password.value)){
    this.authService.login().subscribe(() => {
      if (this.authService.isLoggedIn) {
        this.notification.create('success', '登陆成功', '验证通过');
        // Get the redirect URL from our auth service
        // If no redirect has been set, use the default
        let redirect = this.authService.redirectUrl ? this.router.parseUrl(this.authService.redirectUrl) : '/dashboard';
  
        // Redirect the user
        this.router.navigateByUrl(redirect);
      }
    });
  }
    else{      
      this.notification.create('error', '登陆失败', '验证错误');
  }
     } 
  );
  }
}