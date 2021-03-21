import { Info } from "../info/info";
import { Binary } from "@angular/compiler";

export class Departamentos {
    tipotienda: string
    notienda: Info[]

    constructor(_tipotienda: string, _notienda: Info[]){
        this.tipotienda = _tipotienda
        this.notienda = _notienda

    }

}
