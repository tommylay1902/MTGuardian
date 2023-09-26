import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class PrescriptionService {
  constructor(private http: HttpClient) {}

  getPrescription() {
    return this.http.get('http://localhost:5023/api/v1/prescription');
  }
}
