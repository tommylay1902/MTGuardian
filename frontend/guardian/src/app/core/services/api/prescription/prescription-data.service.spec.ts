import { TestBed } from '@angular/core/testing';

import { PrescriptionDataService } from './prescription-data.service';

describe('PrescriptionDataService', () => {
  let service: PrescriptionDataService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(PrescriptionDataService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
