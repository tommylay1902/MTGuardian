import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RxTabComponent } from './rx-tab.component';

describe('RxTabComponent', () => {
  let component: RxTabComponent;
  let fixture: ComponentFixture<RxTabComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [RxTabComponent]
    });
    fixture = TestBed.createComponent(RxTabComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
