import { Comentarios } from "../Comentarios/comentarios";

export class ComentTienda {
    Tienda: string
    Comentarios: Comentarios[]

    constructor(_Tienda: string, _Comentarios: Comentarios[]){
        this.Tienda = _Tienda
        this.Comentarios = _Comentarios

    }
}
