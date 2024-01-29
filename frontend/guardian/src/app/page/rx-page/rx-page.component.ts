import { Component, OnInit, inject } from '@angular/core';
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { Prescription } from 'src/app/core/models/rx/Prescription';
import { CookieService } from 'ngx-cookie';

import { PrescriptionDataService } from 'src/app/core/services/api/prescription/prescription-data.service';

@Component({
    selector: 'app-rx-page',

    templateUrl: './rx-page.component.html',
    styleUrls: ['./rx-page.component.css'],
})
export class RxPageComponent implements OnInit {
    httpClient = inject(HttpClient);
    cookieService = inject(CookieService);
    prescriptionService = inject(PrescriptionDataService);

    rx: Prescription[] = [];

    ngOnInit(): void {
        this.fetchRx();
    }

    fetchRx() {
        const token = this.cookieService.get('access');
        let headers: HttpHeaders = new HttpHeaders().set(
            'Authorization',
            `Bearer ${token}`
        );

        const params: HttpParams = new HttpParams().set('present', true);
        const options = {
            headers,
            params,
        };

        this.prescriptionService
            .getPrescriptions(
                'http://localhost:8004/api/v1/prescription?present=true',
                options
            )
            .subscribe({
                next: (prescriptions: Prescription[]) => {
                    this.rx = prescriptions;
                },
                error: (error) => {
                    console.error('Error fetching prescriptions:', error);
                },
            });
    }
}
