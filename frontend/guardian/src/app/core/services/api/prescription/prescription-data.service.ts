import {
    HttpClient,
    HttpEvent,
    HttpHeaders,
    HttpParams,
} from '@angular/common/http';
import { Injectable, inject } from '@angular/core';
import { Observable, catchError, map, tap, throwError } from 'rxjs';
import { Prescription } from 'src/app/core/models/rx/Prescription';

@Injectable({
    providedIn: 'root',
})
export class PrescriptionDataService {
    private httpClient = inject(HttpClient);

    getPrescriptions(
        url: string,
        options?: { headers: HttpHeaders; params: HttpParams }
    ): Observable<Prescription[]> {
        return this.httpClient.get<Prescription[]>(url, options);
    }

    //     this.httpClient
    //         .get<Prescription[]>('http://localhost:8004/api/v1/prescription', {
    //             headers: headers,
    //         })
    //         .subscribe({
    //             next: (posts: Prescription[]) => {
    //                 this.rx = posts;
    //             },
    //             error: (err: Error) => {
    //                 return err;
    //             },
    //         });
}
