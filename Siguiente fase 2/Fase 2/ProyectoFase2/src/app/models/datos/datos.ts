import { Departamentos } from "../departamentos/departamentos";
import { Binary } from "@angular/compiler";

export class Datos {
    indice: string
    departinfo: Departamentos[]

    constructor(_indice: string, _departinfo: Departamentos[]){
        this.indice = _indice
        this.departinfo = _departinfo

    }
}
