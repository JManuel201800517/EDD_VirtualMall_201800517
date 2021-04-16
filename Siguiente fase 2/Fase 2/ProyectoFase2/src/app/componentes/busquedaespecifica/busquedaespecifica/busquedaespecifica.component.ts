import { Component, OnInit } from '@angular/core';
import { HttpClient } from "@angular/common/http";
import listadotiendas from "src/json/Tiendas.json";
import { FormControl } from '@angular/forms';
//import { InventarioService } from "../../services/inventario/inventario.service";
import { SubirtiendaService } from 'src/app/services/tienda/subirtienda.service';
import { Tienda } from 'src/app/models/tienda/tienda';
import { Buscarespecifico } from 'src/app/models/buscarespecifico/buscarespecifico';
import { Info } from 'src/app/models/info/info';
import { Guardado } from "src/app/models/Guardado/guardado";
import { ThisReceiver } from '@angular/compiler';


@Component({
  selector: 'app-busquedaespecifica',
  templateUrl: './busquedaespecifica.component.html',
  styleUrls: ['./busquedaespecifica.component.css']
})
export class BusquedaespecificaComponent implements OnInit {

  busqueda = new FormControl('');
  mostrarMensaje = false;
  mostrarMensajeError = false;
  mensajeError = ''

  buscar: any

  //algo!: Guardado;

  constructor(private busquedaService: SubirtiendaService) { }

  ngOnInit(): void {
  }

  subirBusqueda(){
    console.log("Funciona muy bien")
    //console.log(this.inventario.value)
    this.busquedaService.buscarTienda(this.busqueda.value).subscribe((res:any)=>{
      this.mostrarMensaje=true
      this.busqueda.setValue("")
      console.log("Inventario Buscado")
      console.log(res)

      this.buscar = res
      //console.log(this.buscar)

      //this.algo = this.buscar

    },(err)=>{
      this.mostrarMensajeError=true
    })

  }

  desactivarMensaje(){
    this.mostrarMensaje=false
    this.mostrarMensajeError=false
  }

}
