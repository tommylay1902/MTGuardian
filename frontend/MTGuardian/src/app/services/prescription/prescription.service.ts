import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { PrescriptionDTO } from 'src/app/models/prescriptionDTO';

@Injectable({
  providedIn: 'root',
})
export class PrescriptionService {
  constructor(private http: HttpClient) {}

  getPrescription(): Observable<PrescriptionDTO[]> {
    return this.http.get<PrescriptionDTO[]>(
      'http://localhost:5023/api/v1/prescription'
    );
  }
}
