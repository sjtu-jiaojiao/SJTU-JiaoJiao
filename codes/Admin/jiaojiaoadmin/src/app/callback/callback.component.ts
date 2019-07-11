import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { SocialService, JWTTokenModel } from '@delon/auth';
@Component({
  selector: 'app-callback',
  template: ``,
  providers: [SocialService],
})

export class CallbackComponent implements OnInit {
  constructor(
    private socialService: SocialService,
    private route: ActivatedRoute,
  ) {}

  ngOnInit(): void {
    const token = this.route.snapshot.paramMap.get('token');
    this.mockModel(token);
  }

  mockModel(token) {
    const info: any = {
      token: token
    };
    if(typeof(token)==='string')
    this.socialService.callback(info);
  }
}