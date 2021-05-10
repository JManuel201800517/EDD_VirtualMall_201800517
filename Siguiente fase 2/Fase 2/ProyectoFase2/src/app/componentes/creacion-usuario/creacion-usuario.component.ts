import { Component, OnInit } from '@angular/core';
import { HttpClient } from "@angular/common/http";
//import listadotiendas from "src/json/Tiendas.json";
import { FormControl } from '@angular/forms';
import { UsuarioService } from "../../services/usuario/usuario.service";
import { Usuario } from "../../models/usuario/usuario";
import { Cuentas } from "../../models/cuentas/cuentas";

@Component({
  selector: 'app-creacion-usuario',
  templateUrl: './creacion-usuario.component.html',
  styleUrls: ['./creacion-usuario.component.css']
})
export class CreacionUsuarioComponent implements OnInit {

  creacionmasiva = new FormControl('');

  dpi = new FormControl('')
  nombre = new FormControl('')
  cuenta = new FormControl('')
  password = new FormControl('')
  correo = new FormControl('')
  opcion!: string;
  
  mostrarMensaje = false;
  mostrarMensajeError = false;
  mensajeError = ''


  constructor(private usuarioService: UsuarioService) { }

  ngOnInit(): void {
  }

  subirUsuarioMasivo(){
    console.log("Funciona muy bien")
    //console.log(this.inventario.value)


    this.usuarioService.postCuentas(this.creacionmasiva.value).subscribe((res:Cuentas)=>{
      this.mostrarMensaje=true
      this.creacionmasiva.setValue("")
      console.log("Usuarios Masivos Cargados")
      console.log(res)

    }, (err)=>{
      this.mostrarMensajeError=true
    })

    
  }

  subirMerkle(){
    this.usuarioService.getMerkleUsuarios().subscribe((dataList:any)=>{
      console.log(dataList)
      
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError='No se pudo cargar la lista de Inventarios'
    })

  }

  crearUsuario(){


    const usuario: Usuario={
      Dpi:Number(this.dpi.value),
      Nombre:this.nombre.value,
      Correo:this.correo.value,
      Password:this.password.value,
      Cuenta:this.opcion,
    }

    const membresia: Cuentas={
      Usuarios:[usuario]

    }
    
    console.log(this.dpi.value)
    this.usuarioService.postUsuarioIndividual(membresia).subscribe((res:any)=>{
      this.mostrarMensaje=true
      this.dpi.setValue("")
      this.nombre.setValue("")
      this.password.setValue("")
      //this.cui.setValue("")
      this.correo.setValue("")
      console.log("Usuario Creado")
    },(err)=>{
      this.mostrarMensajeError=true
    })


  }


  desactivarMensaje(){
    this.mostrarMensaje=false
    this.mostrarMensajeError=false
  }

}
