import { Injectable, inject } from '@angular/core';
import * as jwt_decode from 'jwt-decode';
import { CookieService } from 'ngx-cookie';
import { JwtTokenService } from './jwt-token.service';
import { HttpClient } from '@angular/common/http';
import { Access } from '../../models/auth/access';
import { Login } from '../../models/auth/login';
import { Router } from '@angular/router';

@Injectable({
    providedIn: 'root',
})
export class AuthService {
    cookieService = inject(CookieService);
    jwtTokenService = inject(JwtTokenService);
    httpClient = inject(HttpClient);
    router = inject(Router);

    public isAuthenticated(): boolean {
        if (this.cookieService.hasKey('access')) {
            const token = this.cookieService.get('access');
            console.log('token', token);
            if (token != undefined) {
                this.jwtTokenService.setToken(token);

                return this.jwtTokenService.isTokenExpired();
            }
        }
        return false;
    }

    public login(req: Login) {
        this.httpClient
            .post<Access>(
                'http://localhost:8004/api/v1/auth/login',
                JSON.stringify(req)
            )
            .subscribe({
                next: (access: Access) => {
                    const now = new Date();
                    now.setHours(now.getHours() + 8);

                    this.cookieService.put('access', access.access);

                    this.router.navigate(['/medication']);
                },
                error: (err: Error) => {
                    return err;
                },
            });
    }
}
