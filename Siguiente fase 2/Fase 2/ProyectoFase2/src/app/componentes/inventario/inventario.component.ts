import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';
import { Comentarios } from 'src/app/models/Comentarios/comentarios';
import { ComentProducto } from 'src/app/models/ComentProducto/coment-producto';
import { Inventario } from 'src/app/models/inventario/inventario';
import { SubComentarios } from 'src/app/models/SubComentarios/sub-comentarios';
import { InventarioService } from "../../services/inventario/inventario.service";

@Component({
  selector: 'app-inventario',
  templateUrl: './inventario.component.html',
  styleUrls: ['./inventario.component.css']
})
export class InventarioComponent implements OnInit {

  lista_estudiantes: number[]=[]

  cargarinventario: any
  //tienda: Tienda

  comentario = new FormControl('');
  subcomentario = new FormControl('');
  nombreInvent = new FormControl('');

  mostrarMensajeError=false
  mostrarMensaje=false
  mensajeError = ''

  opcion!: string;

  constructor(private inventarioService: InventarioService) { }

  ngOnInit(): void {

    this.inventarioService.getListaInventarios().subscribe((dataList:Inventario[])=>{
      console.log(dataList)


      for(var contador=0; contador<dataList.length; contador++ ){
        console.log(dataList[contador])
        this.cargarinventario = dataList[contador]



        var data=dataList[contador].Inventarios
        console.log(data)

        for(var contar=0; contar<data.length; contar++ ){

          var inf = data[contar].Productos
          console.log(inf)

          for(var num=0; num<inf.length; num++){

            console.log(this.cargarinventario)
            //this.nombreInvent.setValue(inf[num].Codigo)

          }


        }

      }
      
      
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError='No se pudo cargar la lista de Inventarios'
    })
    
  }

  desactivarMensaje(){
    this.mostrarMensaje=false
    this.mostrarMensajeError=false
  }

  subirComentario(){

    const comentar: Comentarios={
      Comentario:this.comentario.value,
      SubComentarios:[]
    }

    const producto: ComentProducto={
      Producto:this.nombreInvent.value,
      Comentarios: [comentar]
    }

    console.log(this.comentario.value)
    this.inventarioService.postComentarioProducto(producto).subscribe((res:any)=>{
      this.mostrarMensaje=true
      this.comentario.setValue("")
      this.nombreInvent.setValue("")
      //this.password.setValue("")
      //this.cui.setValue("")
      //this.correo.setValue("")
      console.log("Comentario Subido")
    },(err)=>{
      this.mostrarMensajeError=true
    })

    
  }


  

}
