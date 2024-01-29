import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { RxPageComponent } from './page/rx-page/rx-page.component';
import { AuthComponent } from './page/auth/auth.component';
import { AuthGuardService } from './core/services/auth/auth-guard.service';
import { DashboardComponent } from './page/dashboard/dashboard.component';

const routes: Routes = [
    {
        path: 'auth',
        component: AuthComponent, // Unprotected route
    },
    {
        path: '',
        canActivate: [AuthGuardService], // Protect all child routes
        children: [
            {
                path: '',
                redirectTo: 'dashboard',
                pathMatch: 'full',
            },
            {
                path: 'dashboard',
                component: DashboardComponent, // Protected route
            },
            {
                path: 'medication',
                component: RxPageComponent,
            },
        ],
    },
];

@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule],
})
export class AppRoutingModule {}
