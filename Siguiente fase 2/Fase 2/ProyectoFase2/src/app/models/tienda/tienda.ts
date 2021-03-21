import { Datos } from "../datos/datos";
import { Binary } from "@angular/compiler";

export class Tienda {
    datosinfo: Datos[]

    constructor(_datosinfo: Datos[]){
        this.datosinfo = _datosinfo
    }
}
