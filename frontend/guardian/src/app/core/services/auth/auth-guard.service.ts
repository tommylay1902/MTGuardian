import { Injectable, inject } from '@angular/core';
import { AuthService } from './auth.service';
import { Router } from '@angular/router';

@Injectable({
    providedIn: 'root',
})
export class AuthGuardService {
    authService = inject(AuthService);
    router = inject(Router);

    canActivate(): boolean {
        if (!this.authService.isAuthenticated()) {
            this.router.navigate(['auth']);
            return false;
        }
        return true;
    }
}
