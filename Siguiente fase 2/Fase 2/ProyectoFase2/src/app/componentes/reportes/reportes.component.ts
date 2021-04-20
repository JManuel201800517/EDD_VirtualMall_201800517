import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-reportes',
  templateUrl: './reportes.component.html',
  styleUrls: ['./reportes.component.css']
})
export class ReportesComponent implements OnInit {

  mostrarMensaje = false;
  mostrarMensajeError = false;

  constructor() { }

  ngOnInit(): void {
  }

  CuentasSinCifrar(){

  }

  CuentasCifrar(){

  }

  CuentasCifrarSensible(){
    
  }

  Grafo(){
    
  }

  CaminoCorto(){
    
  }

  Anteriores(){
    
  }


  desactivarMensaje(){
    this.mostrarMensaje=false
    this.mostrarMensajeError=false
  }

}
