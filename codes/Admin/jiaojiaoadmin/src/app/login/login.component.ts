import { Component, OnInit, Inject } from '@angular/core';
import { SocialService } from '@delon/auth';
import { AuthService } from '../auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
  providers: [ SocialService ]

})
export class LoginComponent implements OnInit {

  constructor( public router: Router, private socialService: SocialService , private authService: AuthService) {
  }

  ngOnInit(): void {
  }

  
 login() {
    this.socialService.login('/auth','/callback',{type: 'window'} ).subscribe(res => {
    this.authService.login(res);
    // Redirect the user
    this.router.navigateByUrl(this.router.parseUrl(this.authService.redirectUrl));
    });
    }
  }