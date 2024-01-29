import { Component, Input, OnInit, inject } from '@angular/core';
import {
    Prescription,
    generatePrescriptionTemplate,
} from 'src/app/core/models/rx/Prescription';
import { ModalServiceService } from 'src/app/shared/services/modal-service.service';
import { PrescriptionService } from 'src/app/shared/services/prescription/prescription.service';

@Component({
    selector: 'app-rx-table',
    templateUrl: './rx-table.component.html',
    styleUrls: ['./rx-table.component.css'],
})
export class RxTableComponent implements OnInit {
    @Input() prescriptions: Prescription[] = [];

    modalService = inject(ModalServiceService);
    prescriptionService = inject(PrescriptionService);

    ignoreHeaders: string[] = ['id', 'notes', 'owner'];
    tableHeaders: string[] = Object.keys(generatePrescriptionTemplate()).filter(
        (i) => !this.ignoreHeaders.includes(i)
    );

    ngOnInit(): void {}

    toggleModal(rx: Prescription) {
        this.prescriptionService.setPrescription(rx);
        this.modalService.toggleModal();
    }
}
