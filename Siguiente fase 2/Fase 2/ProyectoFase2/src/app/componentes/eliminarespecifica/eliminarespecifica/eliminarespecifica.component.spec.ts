import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EliminarespecificaComponent } from './eliminarespecifica.component';

describe('EliminarespecificaComponent', () => {
  let component: EliminarespecificaComponent;
  let fixture: ComponentFixture<EliminarespecificaComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ EliminarespecificaComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(EliminarespecificaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
