import { Infopedido } from "../infopedido/infopedido";

export class Pedidos {
    Pedido: Infopedido[]

    constructor(_Pedido: Infopedido[]){
        this.Pedido = _Pedido
    }
}
