import { Component, OnInit } from '@angular/core';
import { HttpClient } from "@angular/common/http";
import listadotiendas from "src/json/Tiendas.json";
import { FormControl } from '@angular/forms';
import { InventarioService } from "../../services/inventario/inventario.service";

@Component({
  selector: 'app-cargar-inventario',
  templateUrl: './cargar-inventario.component.html',
  styleUrls: ['./cargar-inventario.component.css']
})
export class CargarInventarioComponent implements OnInit {

  inventario = new FormControl('');
  mostrarMensaje = false;
  mostrarMensajeError = false;

  constructor(private inventarioservice: InventarioService ) { }

  ngOnInit(): void {
  }
//
  subirInventario(){
    console.log("Funciona muy bien")
    //console.log(this.inventario.value)


    this.inventarioservice.postInventario(this.inventario.value).subscribe((res:any)=>{
      this.mostrarMensaje=true
      this.inventario.setValue("")
      console.log("Inventario Cargado")
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
