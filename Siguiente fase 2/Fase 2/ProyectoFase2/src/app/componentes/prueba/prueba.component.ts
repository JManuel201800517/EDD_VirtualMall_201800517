import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

@Component({
  selector: 'app-prueba',
  templateUrl: './prueba.component.html',
  styleUrls: ['./prueba.component.css']
})
export class PruebaComponent implements OnInit {

  carnet = new FormControl('');
  nombres = new FormControl('');
  apellidos = new FormControl('');
  cui = new FormControl('');
  correo = new FormControl('');
  mostrarMensaje = false;
  mostrarMensajeError = false;

  constructor() { }

  ngOnInit(): void {
  }

  crearEstudiante(){
    console.log("El carnet es: ", this.carnet.value)
  }

}
