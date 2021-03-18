import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CargarTiendasComponent } from './cargar-tiendas.component';

describe('CargarTiendasComponent', () => {
  let component: CargarTiendasComponent;
  let fixture: ComponentFixture<CargarTiendasComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CargarTiendasComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CargarTiendasComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
