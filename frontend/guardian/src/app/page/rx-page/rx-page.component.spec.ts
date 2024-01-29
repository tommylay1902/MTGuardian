import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RxPageComponent } from './rx-page.component';

describe('RxPageComponent', () => {
  let component: RxPageComponent;
  let fixture: ComponentFixture<RxPageComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [RxPageComponent]
    });
    fixture = TestBed.createComponent(RxPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
