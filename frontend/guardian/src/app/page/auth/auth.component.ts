import { HttpClient } from '@angular/common/http';
import {
    AfterViewInit,
    Component,
    ElementRef,
    Renderer2,
    inject,
} from '@angular/core';
import {
    FormBuilder,
    FormControl,
    FormGroup,
    Validators,
} from '@angular/forms';
import { Router } from '@angular/router';
import { CookieService } from 'ngx-cookie';
import { Observable } from 'rxjs';
import { Access } from 'src/app/core/models/auth/access';
import { AuthService } from 'src/app/core/services/auth/auth.service';

@Component({
    selector: 'app-auth',
    templateUrl: './auth.component.html',
    styleUrls: ['./auth.component.css'],
})
export class AuthComponent {
    formBuilder = inject(FormBuilder);
    authService = inject(AuthService);

    loginGroup = this.formBuilder.group({
        email: [
            '',
            Validators.compose([Validators.email, Validators.required]),
        ],
        password: [
            '',
            Validators.compose([Validators.min(14), Validators.required]),
        ],
    });

    onSubmit() {
        if (this.loginGroup.valid) {
            const email = this.loginGroup.get('email')?.value;
            const password = this.loginGroup.get('password')?.value;
            if (
                email != null &&
                email != undefined &&
                password != null &&
                password != undefined
            ) {
                const req = {
                    email,
                    password,
                };

                this.authService.login(req);
            }
        } else {
            alert('invalid');
        }
    }
}
