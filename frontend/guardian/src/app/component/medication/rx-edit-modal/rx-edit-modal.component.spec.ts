import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RxEditModalComponent } from './rx-edit-modal.component';

describe('RxEditModalComponent', () => {
  let component: RxEditModalComponent;
  let fixture: ComponentFixture<RxEditModalComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [RxEditModalComponent]
    });
    fixture = TestBed.createComponent(RxEditModalComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
