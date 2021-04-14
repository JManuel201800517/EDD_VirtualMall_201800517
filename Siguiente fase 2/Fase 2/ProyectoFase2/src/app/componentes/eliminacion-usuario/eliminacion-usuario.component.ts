import { Component, OnInit } from '@angular/core';
import { HttpClient } from "@angular/common/http";
//import listadotiendas from "src/json/Tiendas.json";
import { FormControl } from '@angular/forms'; 
import { UsuarioService } from 'src/app/services/usuario/usuario.service';
import { Eliminarusuario } from "../../models/eliminarusuario/eliminarusuario";

@Component({
  selector: 'app-eliminacion-usuario',
  templateUrl: './eliminacion-usuario.component.html',
  styleUrls: ['./eliminacion-usuario.component.css']
})
export class EliminacionUsuarioComponent implements OnInit {

  //eliminacionusuario = new FormControl('');
  mostrarMensaje = false;
  mostrarMensajeError = false;

  nombre = new FormControl('')
  password = new FormControl('')

  eliminacion = new FormControl("")


  


  constructor(private elimUsuarioService: UsuarioService) { }

  ngOnInit(): void {
  }

  eliminarUsuario(){
    this.eliminacion.setValue('{"Nombre": "' + this.nombre.value +'", "Password": "'+this.password.value+'"}')

    const elimusuario: Eliminarusuario={
      Nombre:this.nombre.value,
      Password:this.password.value,
    }

    console.log("Funciona muy bien")
    //console.log(this.inventario.value)
    console.log(elimusuario)
    console.log(this.eliminacion.value)


    this.elimUsuarioService.deleteUsuario(this.eliminacion.value).subscribe((res:any)=>{
      this.mostrarMensaje=true
      this.nombre.setValue("")
      this.password.setValue("")
      console.log("Usuario Eliminado")
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
