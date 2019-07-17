import { Component, OnInit } from '@angular/core';
import { Location } from '@angular/common';
import {
  FormBuilder,
  FormGroup,
  Validators
} from '@angular/forms';

import { NzNotificationService } from 'ng-zorro-antd';
import { Router } from '@angular/router';
import { WebsiteService } from './../website.service';

@Component({
  selector: 'app-website',
  templateUrl: './website.component.html',
  styleUrls: ['./website.component.css']
})
export class WebsiteComponent implements OnInit {
  validateForm: FormGroup;
  site: any;
  submitForm(): void {
    for (const i in this.validateForm.controls) {
      //if (this.validateForm.controls.hasOwnProperty(i)) {
        this.validateForm.controls[ i ].markAsDirty();
        this.validateForm.controls[ i ].updateValueAndValidity();
      //}
    }
  }
  update(): void{
    this.service.getSite().subscribe( e => {
      this.site =e;
      this.site.name = this.validateForm.controls.name.valid? this.validateForm.controls.name.value: this.site.name;
      this.site.cp = this.validateForm.controls.copyright.valid? this.validateForm.controls.copyright.value: this.site.cp;
      this.site.status = this.validateForm.controls.status.valid? this.validateForm.controls.status.value: this.site.status;
      this.service.updateSite(this.site).subscribe();
     } );
  }

  constructor(
    private service: WebsiteService,
    private location: Location, private router: Router, private fb: FormBuilder, private notification: NzNotificationService) {
  }

  ngOnInit(): void {
    this.validateForm = this.fb.group({
      name: [ null, [ Validators.required ] ],
      copyright: [ null, [ Validators.required ] ],
      status: [ null, [ Validators.required ] ]
    });
  }

}
