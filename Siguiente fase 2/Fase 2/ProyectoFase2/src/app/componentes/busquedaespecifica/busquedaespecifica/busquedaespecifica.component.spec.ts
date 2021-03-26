import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BusquedaespecificaComponent } from './busquedaespecifica.component';

describe('BusquedaespecificaComponent', () => {
  let component: BusquedaespecificaComponent;
  let fixture: ComponentFixture<BusquedaespecificaComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ BusquedaespecificaComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(BusquedaespecificaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
