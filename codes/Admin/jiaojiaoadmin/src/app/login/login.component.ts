import { Component, OnInit } from '@angular/core';
import { Location } from '@angular/common';
import {
  FormBuilder,
  FormGroup,
  Validators
} from '@angular/forms';

import { NzNotificationService } from 'ng-zorro-antd';
import { Router } from '@angular/router';
import { AuthService } from '../auth.service';

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
    public authService: AuthService, public router: Router, 
    private location: Location, private fb: FormBuilder, private notification: NzNotificationService) {
  }

  ngOnInit(): void {
    this.validateForm = this.fb.group({
      userName: [ null, [ Validators.required ] ],
      password: [ null, [ Validators.required ] ]
    });
  }
  login(){
  if(this.validateForm.controls.password.value ==='123456' && this.validateForm.controls.userName.value ==='root'){
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
}