import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable } from 'rxjs';
import {
    Prescription,
    generatePrescriptionTemplate,
} from '../../../core/models/rx/Prescription';

@Injectable({
    providedIn: 'root',
})
export class PrescriptionService {
    private prescriptionSubject: BehaviorSubject<Prescription> =
        new BehaviorSubject<Prescription>(generatePrescriptionTemplate());

    showPrescription$: Observable<Prescription> =
        this.prescriptionSubject.asObservable();

    setPrescription(rx: Prescription) {
        this.prescriptionSubject.next(rx);
    }
}
