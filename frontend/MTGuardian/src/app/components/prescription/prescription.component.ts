import { Component, OnInit } from '@angular/core';
import { PrescriptionDTO } from 'src/app/models/prescriptionDTO';

import { PrescriptionService } from 'src/app/services/prescription/prescription.service';

@Component({
  selector: 'app-prescription',
  templateUrl: './prescription.component.html',
  styleUrls: ['./prescription.component.css'],
})
export class PrescriptionComponent implements OnInit {
  prescriptions: PrescriptionDTO[] = [];

  constructor(private prescription: PrescriptionService) {}

  ngOnInit() {
    this.prescription.getPrescription().subscribe((res) => {
      this.prescriptions = res;
    });
  }
}
