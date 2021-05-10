import { Component, OnInit } from '@angular/core';
import { ComentProducto } from 'src/app/models/ComentProducto/coment-producto';
import { ComentTienda } from 'src/app/models/ComentTienda/coment-tienda';
import {CoTienda} from "../../models/CoTienda/co-tienda";
import {CoProducto} from "../../models/CoProducto/co-producto";
import { InventarioService } from 'src/app/services/inventario/inventario.service';
import { SubirtiendaService } from 'src/app/services/tienda/subirtienda.service';
import { FormControl } from '@angular/forms';
import { SubComentarios } from 'src/app/models/SubComentarios/sub-comentarios';
import { Comentarios } from 'src/app/models/Comentarios/comentarios';

@Component({
  selector: 'app-comentarios',
  templateUrl: './comentarios.component.html',
  styleUrls: ['./comentarios.component.css']
})
export class ComentariosComponent implements OnInit {

  mostrarMensaje = false;
  mostrarMensajeError = false;
  mensajeError = ''

  Cproducto = false;
  Ctienda = false;

  //comentarioProducto: any
  comentarioTienda: any
  //otro: any

  generalProducto: any

  subtienda: any

  Subcomentario = new FormControl("");
  Subcomentario2 = new FormControl("");
  Subcomentario3 = new FormControl("");
  Subcomentario4 = new FormControl("");
  

  constructor(private inventarioService: InventarioService, private tiendaService: SubirtiendaService ) { }

  ngOnInit(): void {
    this.inventarioService.getComentariosProducto().subscribe((dataList:ComentProducto[])=>{
      console.log(dataList)


      for(var contador=0; contador<dataList.length; contador++ ){
        //console.log(dataList[contador])
        //this.comentarioProducto = dataList[contador]
        //this.comentarioProducto = this.comentarioProducto[this.otro]

        //var data=dataList[contador].Inventarios
        //console.log(data)

        const genProducto: CoProducto={
          Datos: dataList
    
        }
        this.generalProducto = genProducto


        const subcomentar: SubComentarios={
          SubComentario:this.Subcomentario.value,
          SubComentarios:[]
        }
        
    
        const coment: Comentarios={
          Comentario: dataList[contador].Comentarios[0].Comentario,
          SubComentarios: [subcomentar]
        }

        const produc: ComentProducto={
          Dpi: dataList[contador].Dpi,
          Producto: dataList[contador].Producto,
          Comentarios:[coment]
        }

        this.subtienda = produc



      }
      

      console.log(this.generalProducto)
      
      
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError='No se pudo cargar la lista de Inventarios'
    })

    this.tiendaService.getComentariosTienda().subscribe((dataList:ComentTienda[])=>{
      console.log(dataList)


      for(var contador=0; contador<dataList.length; contador++ ){
        console.log(dataList[contador])
        //this.comentarioTienda = dataList[contador]

        const genTienda: CoTienda={
          Datos: dataList
        }

        this.comentarioTienda = genTienda

        //var data=dataList[contador].Inventarios
        //console.log(data)

      }

      console.log(this.comentarioTienda)
      
      
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError='No se pudo cargar la lista de Inventarios'
    })
    
  }


  desactivarMensaje(){
    this.mostrarMensaje=false
    this.mostrarMensajeError=false
  }

  ComentariosTienda(){
    this.Ctienda = true
    this.Cproducto = false

  }

  ComentariosProductos(){
    this.Ctienda = false
    this.Cproducto = true
  }

  TablaHashTienda(){

    this.tiendaService.getHashTiendas().subscribe((dataList:any)=>{
      console.log(dataList)
      
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError='No se pudo cargar la Tabla Hash'
    })

  }

  TablaHashProducto(){
    this.inventarioService.getHashProductos().subscribe((dataList:any)=>{
      console.log(dataList)
      
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError='No se pudo cargar la Tabla Hash'
    })

  }



}
