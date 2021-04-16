import { Component, OnInit } from '@angular/core';
import { HttpClient } from "@angular/common/http";
//import listadotiendas from "src/json/Tiendas.json";
import { FormControl } from '@angular/forms'; 
import { LoginService } from 'src/app/services/login/login.service';
import { Eliminarusuario } from 'src/app/models/eliminarusuario/eliminarusuario';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  
  mostrarMensaje = false;
  mostrarMensajeError = false;

  CuentaAdmin = false;
  CuentaUsuario = false;

  usuario = new FormControl('')
  password = new FormControl('')

  sesion = new FormControl("")

  inicio: any

  constructor(private loginService: LoginService) { }

  ngOnInit(): void {
  }

  InicioSesion(){

    this.sesion.setValue('{"Nombre": "' + this.usuario.value +'", "Password": "'+this.password.value+'"}')

    const iniusuario: Eliminarusuario={
      Nombre:this.usuario.value,
      Password:this.password.value,
    }

    //console.log("Funciona muy bien")
    //console.log(this.inventario.value)
    console.log(iniusuario)
    console.log(this.sesion.value)


    console.log("Funciona muy bien")
    //console.log(this.inventario.value)
    this.loginService.inicioSesion(this.sesion.value).subscribe((res:any)=>{
      this.mostrarMensaje=true
      this.sesion.setValue("")
      console.log("Sesion Iniciada")
      console.log(res)

      this.inicio = res
      //console.log(this.buscar)

      //this.algo = this.buscar

      console.log(res.Cuenta)

      if(res.Cuenta == "Usuario"){
        console.log("Es un Usuario")

        this.CuentaUsuario = true
        this.CuentaAdmin = false

      }else if(res.Cuenta == "Admin"){
        console.log("Es un Administrador")

        this.CuentaAdmin = true
        this.CuentaUsuario = false
      }

    },(err)=>{
      this.mostrarMensajeError=true
    })
    
  }





  desactivarMensaje(){
    this.mostrarMensaje=false
    this.mostrarMensajeError=false
  }

}
