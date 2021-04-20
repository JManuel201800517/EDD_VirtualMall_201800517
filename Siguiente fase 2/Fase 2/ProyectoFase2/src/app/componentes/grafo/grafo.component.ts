import { Component, OnInit } from '@angular/core';
import { HttpClient } from "@angular/common/http";
import listadotiendas from "src/json/Tiendas.json";
import { FormControl } from '@angular/forms';
import { GraficasService } from "../../services/graficas/graficas.service";

@Component({
  selector: 'app-grafo',
  templateUrl: './grafo.component.html',
  styleUrls: ['./grafo.component.css']
})
export class GrafoComponent implements OnInit {

  documento = new FormControl('');
  mostrarMensaje = false;
  mostrarMensajeError = false;


  constructor(private Graficaservice: GraficasService) { }

  ngOnInit(): void {
  }

  subirDocumento(){
    console.log("Funciona muy bien")
    //console.log(this.inventario.value)


    this.Graficaservice.postGrafo(this.documento.value).subscribe((res:any)=>{
      this.mostrarMensaje=true
      this.documento.setValue("")
      console.log("Grafo Cargado")
      console.log(res)

    }, (err)=>{
      this.mostrarMensajeError=true
    })

  }

  desactivarMensaje(){
    this.mostrarMensaje=false
    this.mostrarMensajeError=false
  }

}
