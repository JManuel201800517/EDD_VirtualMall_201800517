import { Component, OnInit } from '@angular/core';
import { HttpClient } from "@angular/common/http";
import { SubirtiendaService } from "../../services/tienda/subirtienda.service";
import { Tienda } from 'src/app/models/tienda/tienda';
import { FormControl } from '@angular/forms';
import listadotiendas from "src/json/Tiendas.json";
import { Datos } from 'src/app/models/datos/datos';
import { Departamentos } from 'src/app/models/departamentos/departamentos';
import { Info } from 'src/app/models/info/info';

import * as tiendas from "../../../json/Tiendas.json";


@Component({
  selector: 'app-cargar-tiendas',
  templateUrl: './cargar-tiendas.component.html',
  styleUrls: ['./cargar-tiendas.component.css']
})
export class CargarTiendasComponent implements OnInit {

  Tiendas: any = listadotiendas


  nombre = new FormControl('');

  //nombres = new FormControl('');
  //apellidos = new FormControl('');
  //cui = new FormControl('');
  //correo = new FormControl('');

 // archivo = this.nombre.value+".json"; 



  mostrarMensaje = false;
  mostrarMensajeError = false;

  constructor(private subirtiendaservice: SubirtiendaService) { }

  ngOnInit(): void {
  }

  subirTiendas(){
    console.log("Funciona muy bien")

    //const info: Info={
      //id: Number(this.archivo),
      //nombre: this.archivo,
      //descripcion: this.archivo,
      //contacto: this.archivo,
      //calificacion: Number(this.archivo),
      //logo: this.archivo
    //}

    //const departamentos: Departamentos={
      //tipotienda: this.archivo,
      //notienda:[]    
    //}

    //const datos: Datos={
     // indice: this.archivo,
      //departinfo:[]    
    //}

    //const tienda: Tienda={
      //datosinfo:[]
    //}

    this.subirtiendaservice.postTienda(this.nombre.value).subscribe((res:any)=>{
      this.mostrarMensaje=true
      this.nombre.setValue("")
      console.log("Tiendas Subidas")
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
