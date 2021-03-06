import { Component, OnInit } from '@angular/core';
import { SubirtiendaService } from "../../services/tienda/subirtienda.service";
//import { Curso } from "../../models/curso/curso";
//import { Tienda } from "../../models/tienda/tienda";
import { FormControl } from '@angular/forms';
import { Tienda } from 'src/app/models/tienda/tienda';
import { Datos } from 'src/app/models/datos/datos';
import { Observable } from 'rxjs';
import { Departamentos } from 'src/app/models/departamentos/departamentos';
import { Info } from 'src/app/models/info/info';
import { Comentarios } from 'src/app/models/Comentarios/comentarios';
import { ComentTienda } from 'src/app/models/ComentTienda/coment-tienda';
import { SubComentarios } from 'src/app/models/SubComentarios/sub-comentarios';


@Component({
  selector: 'app-tiendas',
  templateUrl: './tiendas.component.html',
  styleUrls: ['./tiendas.component.css']
})
export class TiendasComponent implements OnInit {

  lista_estudiantes: number[]=[]
  lista_datos: Datos[] = []
  lista_Depa: Departamentos[] = []
  lista_tiendas: Info[] = []

  cargartienda: any
  //tienda: Tienda

  comentario = new FormControl('');
  subcomentario = new FormControl('');
  subcomentario2 = new FormControl('');
  nombreInvent = new FormControl('');
  DPI = new FormControl("");

  mostrarMensajeError=false
  mostrarMensaje=false
  mensajeError = ''

  opcion!: string;

 // datax: string

  //tiendainventario = new FormControl('Hola mundo');

 

  constructor(private subirtiendaService: SubirtiendaService) {


   }

  ngOnInit(): void {

    this.subirtiendaService.getListaTiendas().subscribe((dataList:Tienda[])=>{
      console.log(dataList)


      for(var contador=0; contador<dataList.length; contador++ ){
        console.log(dataList[contador])
        this.cargartienda = dataList[contador]


        var data=dataList[contador].Datos
        console.log(data)

        for(var contar=0; contar<data.length; contar++ ){

          var inf = data[contar].Departamentos
          console.log(inf)

          for(var num=0; num<inf.length; num++){

            console.log(this.cargartienda)

          }


        }

      }
      
      
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError='No se pudo cargar la lista de Tiendas'
    })
    
  }

  desactivarMensaje(){
    this.mostrarMensaje=false
    this.mostrarMensajeError=false
  }

  subirComentario(){
    const subcoment2: SubComentarios={
      SubComentario: this.subcomentario2.value,
      SubComentarios:[]
    }
    
    const subcoment: SubComentarios={
      SubComentario: this.subcomentario.value,
      SubComentarios:[subcoment2]
    }

    const comentar: Comentarios={
      Comentario:this.comentario.value,
      SubComentarios:[subcoment]
    }

    const tienda: ComentTienda={
      Dpi: Number(this.DPI.value),
      Tienda:this.nombreInvent.value,
      Comentarios: [comentar]
    }

    console.log(this.comentario.value)
    this.subirtiendaService.postComentarioTienda(tienda).subscribe((res:any)=>{
      this.mostrarMensaje=true
      this.comentario.setValue("")
      this.nombreInvent.setValue("")
      this.subcomentario.setValue("")
      this.DPI.setValue("")
      this.subcomentario2.setValue("")
      //this.password.setValue("")
      //this.cui.setValue("")
      //this.correo.setValue("")
      console.log("Comentario Subido")
    },(err)=>{
      this.mostrarMensajeError=true
    })

    
  }

}
