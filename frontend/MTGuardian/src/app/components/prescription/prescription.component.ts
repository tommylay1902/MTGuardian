import { Component } from '@angular/core';

import { PrescriptionService } from 'src/app/services/prescription/prescription.service';

@Component({
  selector: 'app-prescription',
  templateUrl: './prescription.component.html',
  styleUrls: ['./prescription.component.css'],
})
export class PrescriptionComponent {
  prescriptions: [] = [];

  constructor(private prescription: PrescriptionService) {}
  ngOnInit() {
    this.prescription.getPrescription().subscribe((res) => console.log(res));
  }
}
