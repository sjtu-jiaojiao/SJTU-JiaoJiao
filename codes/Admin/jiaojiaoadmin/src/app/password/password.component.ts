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
import { AdminService } from '../admin.service';
@Component({
  selector: 'app-password',
  templateUrl: './password.component.html',
  styleUrls: ['./password.component.css']
})
export class PasswordComponent implements OnInit {

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
    private adminService: AdminService,
     private fb: FormBuilder, private notification: NzNotificationService) {
  }

  ngOnInit(): void {
    this.validateForm = this.fb.group({
      userName: [ null, [ Validators.required ] ],
      password: [ null, [ Validators.required ] ],
      newpassword: [ null, [ Validators.required ] ]
    });
  }

  editpwd() {
  this.adminService.editpwd
    (this.validateForm.controls.userName.value , this.validateForm.controls.password.value , 
      this.validateForm.controls.newpassword.value);
  }
}
