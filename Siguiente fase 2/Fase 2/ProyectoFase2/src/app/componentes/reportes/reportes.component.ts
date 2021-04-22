import { Component, OnInit } from '@angular/core';
import { GraficasService } from "../../services/graficas/graficas.service";
import { HttpClient } from "@angular/common/http";

@Component({
  selector: 'app-reportes',
  templateUrl: './reportes.component.html',
  styleUrls: ['./reportes.component.css']
})
export class ReportesComponent implements OnInit {

  mostrarMensaje = false;
  mostrarMensajeError = false;
  mensajeError = ''

  constructor(private graficaService: GraficasService) { }

  ngOnInit(): void {
  }

  CuentasSinCifrar(){
    this.graficaService.getArbolSinCifrar().subscribe((dataList:any)=>{
      console.log(dataList)
      
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError='No se pudo cargar la lista de Inventarios'
    })

  }

  CuentasCifrar(){
    this.graficaService.getArbolCifrar().subscribe((dataList:any)=>{
      console.log(dataList)
      
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError='No se pudo cargar la lista de Inventarios'
    })

  }

  CuentasCifrarSensible(){
    this.graficaService.getArbolCifrarSuave().subscribe((dataList:any)=>{
      console.log(dataList)
      
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError='No se pudo cargar la lista de Inventarios'
    })
    
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
