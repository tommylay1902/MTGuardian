import {
  AfterViewInit,
  ChangeDetectorRef,
  Component,
  OnInit,
} from '@angular/core';
import { PrescriptionDTO } from 'src/app/models/prescriptionDTO';

import { PrescriptionService } from 'src/app/services/prescription/prescription.service';

interface Column {
  field: string;
  header: string;
}

@Component({
  selector: 'app-prescription',
  templateUrl: './prescription.component.html',
  styleUrls: ['./prescription.component.css'],
})
export class PrescriptionComponent implements OnInit, AfterViewInit {
  prescriptions: PrescriptionDTO[] = [];
  cols!: Column[];
  dataLoaded: boolean = false;

  constructor(
    private prescription: PrescriptionService,
    private cdr: ChangeDetectorRef
  ) {}

  ngOnInit() {
    this.loadData();
  }

  ngAfterViewInit() {}

  private setColumns() {
    if (this.prescriptions.length > 0) {
      const keys = Object.keys(this.prescriptions[0]);

      this.cols = keys.map((key) => ({
        field: key,
        header: key.charAt(0).toUpperCase() + key.slice(1), // Capitalize the first letter of the key
      }));
    } else {
      this.cols = [];
    }
  }

  private loadData() {
    setTimeout(() => {});
    this.prescription.getPrescription().subscribe((result) => {
      this.prescriptions = result;
      this.setColumns();
      this.dataLoaded = true;
    });
  }
}
