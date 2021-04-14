import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EliminacionUsuarioComponent } from './eliminacion-usuario.component';

describe('EliminacionUsuarioComponent', () => {
  let component: EliminacionUsuarioComponent;
  let fixture: ComponentFixture<EliminacionUsuarioComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ EliminacionUsuarioComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(EliminacionUsuarioComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
