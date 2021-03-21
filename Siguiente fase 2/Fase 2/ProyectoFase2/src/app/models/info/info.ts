import { Binary } from "@angular/compiler";

export class Info {
    id: number
    nombre: string
    descripcion: string
    contacto: string
    calificacion: number
    logo: string

    constructor(_id: number, _nombre: string, _descripcion: string, _contacto: string, _calificacion: number, _logo: string){
        this.id = _id
        this.nombre = _nombre
        this.descripcion = _descripcion
        this.contacto = _contacto
        this.calificacion = _calificacion
        this.logo = _logo
    }

}
