import { Injectable } from '@angular/core';
import * as jwt from 'jwt-decode';

@Injectable()
export class JwtTokenService {
    jwtToken: string = '';
    decodedToken: { [key: string]: string } = {};

    setToken(token: string) {
        if (token) {
            this.jwtToken = token;
        }
    }

    decodeToken() {
        if (this.jwtToken) {
            this.decodedToken = jwt.jwtDecode(this.jwtToken);
        }
    }

    getDecodeToken() {
        return jwt.jwtDecode(this.jwtToken);
    }

    getEmailId() {
        this.decodeToken();
        return this.decodedToken ? this.decodedToken['sub'] : null;
    }

    getExpiryTime() {
        this.decodeToken();
        return this.decodedToken ? this.decodedToken['exp'] : null;
    }

    isTokenExpired(): boolean {
        const expiryTime: number | null =
            this.getExpiryTime() !== null ? Number(this.getExpiryTime()) : null;
        if (expiryTime !== null) {
            return 1000 * expiryTime - new Date().getTime() > 0;
        } else {
            console.log('we running');
            return false;
        }
    }
}
