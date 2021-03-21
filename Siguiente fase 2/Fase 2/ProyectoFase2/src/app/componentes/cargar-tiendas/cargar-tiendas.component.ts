import { Component, OnInit } from '@angular/core';
import { HttpClient } from "@angular/common/http";
import { SubirtiendaService } from "../../services/tienda/subirtienda.service";
import { Tienda } from 'src/app/models/tienda/tienda';
import { FormControl } from '@angular/forms';
import listadotiendas from "src/json/Tiendas.json";


@Component({
  selector: 'app-cargar-tiendas',
  templateUrl: './cargar-tiendas.component.html',
  styleUrls: ['./cargar-tiendas.component.css']
})
export class CargarTiendasComponent implements OnInit {

  Tiendas: any = listadotiendas

 // carnet = new FormControl('');
  //nombres = new FormControl('');
  //apellidos = new FormControl('');
  //cui = new FormControl('');
  //correo = new FormControl('');
  mostrarMensaje = false;
  mostrarMensajeError = false;

  constructor(private subirtiendaservice: SubirtiendaService) { }

  ngOnInit(): void {
  }

  subirTiendas(){
    console.log("Funciona muy bien")
    //const tienda: Tienda={
      //datosinfo:[]
    //}

    //this.subirtiendaservice.postTienda(tienda).subscribe((res:any)=>{
      //this.mostrarMensaje=true

    //}, (err)=>{
     // this.mostrarMensajeError=true
    //})
  }

}
