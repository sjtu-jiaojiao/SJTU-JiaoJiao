import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { DashboardComponent } from './dashboard/dashboard.component';
import { UserDetailComponent } from './userdetail/userdetail.component';
import { UserComponent } from './user/user.component';
import { LoginComponent } from './login/login.component';
import { InfoComponent } from './info/info.component';
import { InfoDetailComponent } from './infodetail/infodetail.component';
import { StatisticComponent } from './statistic/statistic.component';

const routes: Routes = [
  { path: 'dashboard', component: DashboardComponent },
  { path: '', redirectTo: '/dashboard', pathMatch: 'full' },
  { path: 'user/:id', component: UserDetailComponent },
  { path: 'user', component: UserComponent},
  { path: 'login', component: LoginComponent},
  { path: 'info', component:  InfoComponent},
  { path: 'info/:id', component:  InfoDetailComponent},
  { path: 'statistic', component:  StatisticComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
