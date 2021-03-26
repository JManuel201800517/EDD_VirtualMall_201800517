import { Component, OnInit } from '@angular/core';
import { Inventario } from 'src/app/models/inventario/inventario';
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


  

}
