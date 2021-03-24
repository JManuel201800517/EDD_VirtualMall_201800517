import { Pedirproductos } from "../pedirproductos/pedirproductos";

export class Infopedido {
    FechaPedido: string
    TiendaPedido: string
    DepartamentoPedido: string
    CalificacionPedido: number
    ProductosPedido: Pedirproductos[]

    constructor(_FechaPedido: string, _TiendaPedido: string, _DepartamentoPedido: string,
        _CalificacionPedido: number, _ProductosPedido: Pedirproductos[]){
        this.FechaPedido = _FechaPedido
        this.TiendaPedido = _TiendaPedido
        this.DepartamentoPedido = _DepartamentoPedido
        this.CalificacionPedido = _CalificacionPedido
        this.ProductosPedido = _ProductosPedido

    }
}
