import { Infoinventario } from "../infoinventario/infoinventario";

export class Inventario {
    DatosInventario: Infoinventario[]

    constructor(_DatosInventario: Infoinventario[]){
        this.DatosInventario = _DatosInventario
    }
    
}
