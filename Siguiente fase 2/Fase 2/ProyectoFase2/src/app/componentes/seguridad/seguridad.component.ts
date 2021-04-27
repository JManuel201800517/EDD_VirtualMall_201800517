import { Component, OnInit } from '@angular/core';
import { HttpClient } from "@angular/common/http";
import listadotiendas from "src/json/Tiendas.json";
import { FormControl } from '@angular/forms';
import { UsuarioService } from "../../services/usuario/usuario.service";
import { Cuentas } from 'src/app/models/cuentas/cuentas';

@Component({
  selector: 'app-seguridad',
  templateUrl: './seguridad.component.html',
  styleUrls: ['./seguridad.component.css']
})
export class SeguridadComponent implements OnInit {

  clave = new FormControl('');
  mostrarMensaje = false;
  mostrarMensajeError = false;
  mostrarBoton = false;
  mensajeError = ''

  cargarusuario: any


  constructor(private usuarioService: UsuarioService) { }

  ngOnInit(): void {
  }

  subirClave(){
    if (this.clave.value == "JoyBoy-Raftel"){
      this.mostrarMensaje=true
      this.mostrarBoton=true

    }else{
      this.mostrarMensajeError=true
      this.clave.setValue("")
      this.mensajeError='No se pudo cargar la lista de Inventarios'
    }

  }

  mostrarNormal(){
    this.usuarioService.getListaUsuarios().subscribe((dataList:Cuentas[])=>{
      console.log(dataList)


      for(var contador=0; contador<dataList.length; contador++ ){
        console.log(dataList[contador])
        this.cargarusuario = dataList[contador]


        var data=dataList[contador].Usuarios
        console.log(data)

        for(var contar=0; contar<data.length; contar++ ){

          var inf = data[contar].Nombre
          console.log(inf)

          for(var num=0; num<inf.length; num++){

            console.log(this.cargarusuario)

          }


        }

      }
      
      
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError='No se pudo cargar la lista de Inventarios'
    })

  }

  mostrarEncriptado(){
    console.log("Funciona muy bien")
    //console.log(this.inventario.value)


    this.usuarioService.getUsuariosEncriptados().subscribe((dataList:any)=>{
      console.log(dataList)
      
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
