import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NavbarComponent } from './component/navbar/navbar.component';

import { RxPageComponent } from './page/rx-page/rx-page.component';
import { RxTableComponent } from './component/medication/rx-table/rx-table.component';
import { RxTabComponent } from './component/medication/rx-tab/rx-tab.component';
import { HttpClientModule } from '@angular/common/http';
import { AuthComponent } from './page/auth/auth.component';
import { CookieModule } from 'ngx-cookie';
import { JwtTokenService } from './core/services/auth/jwt-token.service';
import { DashboardComponent } from './page/dashboard/dashboard.component';
import { ReactiveFormsModule } from '@angular/forms';
import { RxEditModalComponent } from './component/medication/rx-edit-modal/rx-edit-modal.component';

@NgModule({
    declarations: [
        AppComponent,
        NavbarComponent,
        RxPageComponent,
        RxTableComponent,
        RxTabComponent,
        AuthComponent,
        DashboardComponent,
        RxEditModalComponent,
    ],
    imports: [
        BrowserModule,
        AppRoutingModule,
        HttpClientModule,
        ReactiveFormsModule,
        CookieModule.withOptions(),
    ],
    providers: [JwtTokenService],
    bootstrap: [AppComponent],
})
export class AppModule {}
