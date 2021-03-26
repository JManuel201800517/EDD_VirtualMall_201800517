import { Component, OnInit } from '@angular/core';
import { Pedidos } from 'src/app/models/pedidos/pedidos';
import { PedidosService } from "../../services/pedidos/pedidos.service";

@Component({
  selector: 'app-calendario-pedidos',
  templateUrl: './calendario-pedidos.component.html',
  styleUrls: ['./calendario-pedidos.component.css']
})
export class CalendarioPedidosComponent implements OnInit {

  lista_estudiantes: number[]=[]

  cargarpedido: any
  //tienda: Tienda

  mostrarMensajeError=false
  mostrarMensaje=false
  mensajeError = ''

  opcion!: string;

  constructor(private pedidosService: PedidosService) { }

  ngOnInit(): void {
    this.pedidosService.getListaPedidos().subscribe((dataList:Pedidos[])=>{
      console.log(dataList)


      for(var contador=0; contador<dataList.length; contador++ ){
        console.log(dataList[contador])
        this.cargarpedido = dataList[contador]


        var data=dataList[contador].Pedidos
        console.log(data)

        for(var contar=0; contar<data.length; contar++ ){

          var inf = data[contar].Productos
          console.log(inf)

          for(var num=0; num<inf.length; num++){

            console.log(this.cargarpedido)

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
