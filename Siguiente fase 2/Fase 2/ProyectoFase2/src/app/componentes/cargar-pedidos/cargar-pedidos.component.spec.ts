import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CargarPedidosComponent } from './cargar-pedidos.component';

describe('CargarPedidosComponent', () => {
  let component: CargarPedidosComponent;
  let fixture: ComponentFixture<CargarPedidosComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CargarPedidosComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CargarPedidosComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
