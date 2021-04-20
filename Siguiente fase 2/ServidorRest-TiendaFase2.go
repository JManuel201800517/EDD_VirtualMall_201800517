package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

var cargartienda []Tienda

type Tienda struct {
	DATOSINFO []DATOS `json:"Datos,omitempty"`
}

type DATOS struct {
	INDICE     string          `json:"Indice,omitempty"`
	DEPARTINFO []DEPARTAMENTOS `json:"Departamentos,omitempty"`
}

type DEPARTAMENTOS struct {
	TIPOTIENDA string `json:"Nombre,omitempty"`
	NOTIENDA   []info `json:"Tiendas,omitempty"`
}

type info struct {
	ID           string `json:"Id,omitempty"`
	NOMBRE       string `json:"Nombre,omitempty"`
	DESCRIPCION  string `json:"Descripcion,omitempty"`
	CONTACTO     string `json:"Contacto,omitempty"`
	CALIFICACION int    `json:"Calificacion,omitempty"`
	LOGO         string `json:"Logo,omitempty"`
}

var tiendaespecifica []Especifica

type Especifica struct {
	DepartamentoEspecifico string `json:"Departamento,omitempty"`
	NombreEspecifico       string `json:"Nombre,omitempty"`
	CalificacionEspecifico int    `json:"Calificacion,omitempty"`
}

var eliminar []Eliminacion

type Eliminacion struct {
	NombreEliminar       string `json:"Nombre,omitempty"`
	CategoriaEliminar    string `json:"Categoria,omitempty"`
	CalificacionEliminar int    `json:"Calificacion,omitempty"`
}

var invent []Inventario

type Inventario struct{
	DatosInventario  []InfoInventario  `json:"Inventarios,omitempty"`
}

type InfoInventario struct{
	InvenTienda          string      `json:"Tienda,omitempty"`
	InvenDepartamento    string      `json:"Departamento,omitempty"`
	InvenCalificacion    int         `json:"Calificacion,omitempty"`
	ProductosInventario  []Productos `json:"Productos,omitempty"`
}

type Productos struct{
	NombreProducto      string  `json:"Nombre,omitempty"`
	CodigoProducto      int     `json:"Codigo,omitempty"`
	DescripcionProducto string  `json:"Descripcion,omitempty"`
	PrecioProducto      float64 `json:"Precio,omitempty"`
	CantidadProducto    int     `json:"Cantidad,omitempty"`
	ImagenProducto      string  `json:"Imagen,omitempty"`
	AlmacenProducto     string  `json:"Almacenamiento,omitempty"`
}


var pedir []Pedidos

type Pedidos struct{
	PEDIDO  []InfoPedido  `json:"Pedidos,omitempty"`
}

type InfoPedido struct{
	FechaPedido         string            `json:"Fecha,omitempty"`
	TiendaPedido        string            `json:"Tienda,omitempty"`
	DepartamentoPedido  string            `json:"Departamento,omitempty"`
	CalificacionPedido  int               `json:"Calificacion,omitempty"`
	ClientePedido       int               `json:"Cliente,omitempty"` 
	ProductosPedido     []PedirProductos  `json:"Productos,omitempty"`
}

type PedirProductos struct{
	ProductoCodigo  int  `json:"Codigo,omitempty"`     
}


var miembro []Cuentas

type Cuentas struct{
	USUARIOS  []Usuarios  `json:"Usuarios,omitempty"`
}

type Usuarios struct{
	Dpi        int       `json:"Dpi,omitempty"`
	Nombre     string    `json:"Nombre,omitempty"`
	Correo     string    `json:"Correo,omitempty"`
	Password   string    `json:"Password,omitempty"`
	Cuenta     string    `json:"Cuenta,omitempty"`

}

var suprimir []EliminarUsuario

type EliminarUsuario struct {
	NombreUsuario       string `json:"Nombre,omitempty"`
	PasswordUsuario     string `json:"Password,omitempty"`
}


var ingresar []Sesion

type Sesion struct {
	NombreUsuario       string `json:"Nombre,omitempty"`
	PasswordUsuario     string `json:"Password,omitempty"`
}


var caminografo []Grafo

type Grafo struct {
	Nodos          []Nodo  `json:"Nodos,omitempty"`
	PosicionRobot  string  `json:"PosicionInicialRobot,omitempty"`
	Entrega        string  `json:"Entrega,omitempty"`
}

type Nodo struct {
	Nombre    string     `json:"Nombre,omitempty"`
	Enlace    []Enlaces  `json:"Enlaces,omitempty"`
}

type Enlaces struct {
	Nombre     string   `json:"Nombre,omitempty"`
	Distancia  int      `json:"Distancia,omitempty"`
}


// no usar esta seccion encacillada{

var lista []infolista

type infolista struct {
	ID           string `json:"Id,omitempty"`
	NOMBRE       string `json:"Nombre,omitempty"`
	DESCRIPCION  string `json:"Descripcion,omitempty"`
	CONTACTO     string `json:"Contacto,omitempty"`
	CALIFICACION int    `json:"Calificacion,omitempty"`
	Next         *infolista
}

func Ordenamiento(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//json.NewEncoder(w).Encode(cargartienda)

	var linea Tienda

	//var multi Linealizacion

	//z := 0

	p1 := len(linea.DATOSINFO)
	p2 := len(linea.DATOSINFO[0].DEPARTINFO)

	// k = Numero de Calificaciones

	k := 5

	multi := p1 * p2 * k

	///lista := list.New()

	vector := make([]infolista, multi)

	z := 0
	for z < len(linea.DATOSINFO) {

		j := 0

		for j < len(linea.DATOSINFO[z].DEPARTINFO) {

			i := 0

			for i < len(linea.DATOSINFO[z].DEPARTINFO[j].NOTIENDA) {
				p3 := linea.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i].CALIFICACION

				//nom := linea.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i].NOMBRE

				n := linea.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i].NOMBRE
				d := linea.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i].DESCRIPCION
				c := linea.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i].CONTACTO

				vectorlinea := (((z * p2) + j) * k) + p3

				//vec = append(vec, Linealizacion{NOMBRE: n, DESCRIPCION: d, CONTACTO: c, CALIFICACION: p3})
				//x := lista.PushBack(nom)
				vector[vectorlinea].Next.NOMBRE = n
				vector[vectorlinea].Next.DESCRIPCION = d
				vector[vectorlinea].Next.CONTACTO = c
				vector[vectorlinea].Next.CALIFICACION = p3

				//vector[vectorlinea] = vec

				//json.NewEncoder(w).Encode(vec)

				//vector[vectorlinea] = item.DATOSINFO[z].DEPARTINFO[j].TIPOTIENDA[i].NOMBRE[vectorlinea]
				i = i + 1

			}

			j = j + 1

		}

		z = z + 1

	}

	json.NewEncoder(w).Encode(lista)

}


// } no usar esta seccion encillada



func ListadoDeTiendasEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cargartienda)

}



func ListadoDeUsuariosEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(miembro)

}

func ObtenerTipoCuentaEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for _, item := range miembro {
		for z := 0; z < len(item.USUARIOS); z++ {
			if item.USUARIOS[z].Cuenta == params["Cuenta"] {
				json.NewEncoder(w).Encode(item.USUARIOS[z])
				return
			}

		}
	}
	json.NewEncoder(w).Encode(&Cuentas{})
}

func CrearUsuarioEndPoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var avatar Cuentas
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "Insertar dato valido")
	}

	json.Unmarshal(reqBody, &avatar)

	for z := 0; z < len(avatar.USUARIOS); z++ {
		_ = json.NewDecoder(req.Body).Decode(&avatar)
		
	}
	miembro = append(miembro, avatar)
	json.NewEncoder(w).Encode(avatar)

}

func EliminarUsuarioEndPoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(req)
	var elimus EliminarUsuario
	json.NewDecoder(req.Body).Decode(&elimus)

	//tienda.ID = params["Id nuevo"]
	suprimir = append(suprimir, elimus)
	//json.NewEncoder(w).Encode(tiendaespecifica)

	for _, item := range miembro {
		z := 0

		for  z < len(item.USUARIOS) {

			if item.USUARIOS[z].Nombre == elimus.NombreUsuario{

				if item.USUARIOS[z].Password == elimus.PasswordUsuario{

					item.USUARIOS = append(item.USUARIOS[:z], item.USUARIOS[z+1:]...)

					break
				}


				break
			}

			z = z + 1


		}

	}

	json.NewEncoder(w).Encode(miembro)

}


func SubirGrafoEndPoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var robot Grafo
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "Insertar dato valido")
	}

	json.Unmarshal(reqBody, &robot)

	for z := 0; z < len(robot.Nodos); z++ {
		for j := 0; j < len(robot.Nodos[z].Enlace); j++ {
				_ = json.NewDecoder(req.Body).Decode(&robot)
				//tienda.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i].ID = strconv.Itoa(len(cargartienda) + 1)

		}
	}
	caminografo = append(caminografo, robot)
	json.NewEncoder(w).Encode(robot)

}


func DatosGrafoEndPoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(caminografo)

}





func ListadoDeInventarioEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invent)

}

func ListadoDePedidoEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pedir)

}

func ObtenerIndiceTiendaEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for _, item := range cargartienda {
		for z := 0; z < len(item.DATOSINFO); z++ {
			if item.DATOSINFO[z].INDICE == params["Indice"] {
				json.NewEncoder(w).Encode(item.DATOSINFO[z])
				return
			}

		}
	}
	json.NewEncoder(w).Encode(&Tienda{})
}

func ObtenerNombreTipoTiendaEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	for _, item := range cargartienda {
		for z := 0; z < len(item.DATOSINFO); z++ {

			j := 0

			for j < len(item.DATOSINFO[z].DEPARTINFO) {
				if item.DATOSINFO[z].DEPARTINFO[j].TIPOTIENDA == params["Nombre"] {
					json.NewEncoder(w).Encode(item.DATOSINFO[z].DEPARTINFO[j])
					return
				}
				j = j + 1
			}

		}
	}
	json.NewEncoder(w).Encode(&Tienda{})
}

func ObtenerNombreTiendaEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	for _, item := range cargartienda {
		for z := 0; z < len(item.DATOSINFO); z++ {

			j := 0

			for j < len(item.DATOSINFO[z].DEPARTINFO) {

				i := 0

				for i < len(item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA) {
					if item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i].NOMBRE == params["Nombre"] {
						json.NewEncoder(w).Encode(item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i])

						return
					}

					i = i + 1

				}

				j = j + 1

			}

		}
	}
	json.NewEncoder(w).Encode(&Tienda{})
}

func ObtenerContactoTiendaEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	for _, item := range cargartienda {
		for z := 0; z < len(item.DATOSINFO); z++ {

			j := 0

			for j < len(item.DATOSINFO[z].DEPARTINFO) {

				i := 0

				for i < len(item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA) {
					if item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i].CONTACTO == params["Contacto"] {
						json.NewEncoder(w).Encode(item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i])

						return
					}

					i = i + 1

				}

				j = j + 1

			}

		}
	}
	json.NewEncoder(w).Encode(&Tienda{})
}

func BorrarTiendaEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for _, item := range cargartienda {

		z := 0
		for z < len(item.DATOSINFO) {

			j := 0

			for j < len(item.DATOSINFO[z].DEPARTINFO) {

				i := 0

				for i < len(item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA) {
					if item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i].NOMBRE == params["Nombre"] {
						item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA = append(item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[:i], item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i+1:]...)

						//_ = json.NewDecoder(req.Body).Decode(&index)
						break
					}

					i = i + 1

				}

				j = j + 1

			}

			z = z + 1

		}

	}
	json.NewEncoder(w).Encode(cargartienda)

}

func CrearEndPoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tienda Tienda
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "Insertar dato valido")
	}

	json.Unmarshal(reqBody, &tienda)

	for z := 0; z < len(tienda.DATOSINFO); z++ {
		for j := 0; j < len(tienda.DATOSINFO[z].DEPARTINFO); j++ {
			for i := 0; i < len(tienda.DATOSINFO[z].DEPARTINFO[j].NOTIENDA); i++ {
				_ = json.NewDecoder(req.Body).Decode(&tienda)
				//tienda.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i].ID = strconv.Itoa(len(cargartienda) + 1)
			}

		}
	}
	cargartienda = append(cargartienda, tienda)
	json.NewEncoder(w).Encode(tienda)

}

func TiendaEspecificaEndPoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(req)
	var especifico Especifica
	json.NewDecoder(req.Body).Decode(&especifico)

	//tienda.ID = params["Id nuevo"]
	tiendaespecifica = append(tiendaespecifica, especifico)
	//json.NewEncoder(w).Encode(tiendaespecifica)

	for _, item := range cargartienda {
		for z := 0; z < len(item.DATOSINFO); z++ {

			j := 0

			for j < len(item.DATOSINFO[z].DEPARTINFO) {

				if item.DATOSINFO[z].DEPARTINFO[j].TIPOTIENDA == especifico.DepartamentoEspecifico {

					i := 0

					for i < len(item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA) {

						if item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i].NOMBRE == especifico.NombreEspecifico {

							if item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i].CALIFICACION == especifico.CalificacionEspecifico {

								json.NewEncoder(w).Encode(item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i])
								//z = len(item.DATOSINFO)
								//j = len(item.DATOSINFO[z].DEPARTINFO)
								//i = len(item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA)

								return
							}

							return
						}

						i = i + 1

					}

					return
				}

				j = j + 1

			}

		}

	}

	json.NewEncoder(w).Encode(&Tienda{})

}

func EliminarRegistroEndPoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(req)
	var elim Eliminacion
	json.NewDecoder(req.Body).Decode(&elim)

	//tienda.ID = params["Id nuevo"]
	eliminar = append(eliminar, elim)
	//json.NewEncoder(w).Encode(tiendaespecifica)

	for _, item := range cargartienda {
		for z := 0; z < len(item.DATOSINFO); z++ {

			j := 0

			for j < len(item.DATOSINFO[z].DEPARTINFO) {

				if item.DATOSINFO[z].DEPARTINFO[j].TIPOTIENDA == elim.CategoriaEliminar {

					i := 0

					for i < len(item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA) {

						if item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i].NOMBRE == elim.NombreEliminar {

							if item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i].CALIFICACION == elim.CalificacionEliminar {

								item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA = append(item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[:i], item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i+1:]...)
								//z = len(item.DATOSINFO)
								//j = len(item.DATOSINFO[z].DEPARTINFO)
								//i = len(item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA)

								break
							}

							break
						}

						i = i + 1

					}

					break
				}

				j = j + 1

			}

		}

	}

	json.NewEncoder(w).Encode(cargartienda)

}

func InventarioEndPoint(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var produc Inventario
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "Insertar dato valido")
	}

	json.Unmarshal(reqBody, &produc)

	for z := 0; z < len(produc.DatosInventario); z++ {
		for j := 0; j < len(produc.DatosInventario[z].ProductosInventario); j++ {
			_ = json.NewDecoder(req.Body).Decode(&produc)

		}
	}
	invent = append(invent, produc)
	json.NewEncoder(w).Encode(produc)


}

func PedidosEndPoint(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var peticion Pedidos
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "Insertar dato valido")
	}

	json.Unmarshal(reqBody, &peticion)

	for z := 0; z < len(peticion.PEDIDO); z++ {
		for j := 0; j < len(peticion.PEDIDO[z].ProductosPedido); j++ {
			_ = json.NewDecoder(req.Body).Decode(&peticion)

		}
	}
	pedir = append(pedir, peticion)
	json.NewEncoder(w).Encode(peticion)

	//for _, item := range invent {

	//	for _, otro := range pedir{

	//		for z := 0; z < len(item.DatosInventario); z++ {

	//			j := 0
	
	//			for j < len(item.DatosInventario[z].TiendaPedido) {
	
	//				q := 0


	
					//if item.DATOSINFO[z].DEPARTINFO[j].TIPOTIENDA == peticion.Pedidos {
	
					//	i := 0
	
						
	
					//	return
					//}
	
	//				j = j + 1
	
	//			}
	
	//		}


	//	}


	//}

	//json.NewEncoder(w).Encode(&Inventario{})
	
}


func InicioSesionEndPoint(w http.ResponseWriter, req *http.Request){

	w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(req)
	var usuario Sesion
	json.NewDecoder(req.Body).Decode(&usuario)

	//tienda.ID = params["Id nuevo"]
	ingresar = append(ingresar, usuario)
	//json.NewEncoder(w).Encode(tiendaespecifica)

	for _, item := range miembro {
		for z := 0; z < len(item.USUARIOS); z++ {

			if item.USUARIOS[z].Nombre == usuario.NombreUsuario {

				if item.USUARIOS[z].Password == usuario.PasswordUsuario {

					json.NewEncoder(w).Encode(item.USUARIOS[z])
					//z = len(item.DATOSINFO)
					//j = len(item.DATOSINFO[z].DEPARTINFO)
					//i = len(item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA)

					return
				}

				return
			}

		}

	}

	json.NewEncoder(w).Encode(&Cuentas{})
	
	
}



// no usar esta seccion encacillada{


func DescargarArchivoEndPoint(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Disposition", "attachment; filename=prueba1.0")
	w.Header().Set("Content-Type", "application/json")

}

func Grafica(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cargartienda)

	g := newGraph()
	// https://graphviz.gitlab.io/_pages/Gallery/directed/fsm.html

	//w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(req)
	for _, item := range cargartienda {

		z := 0
		for z < len(item.DATOSINFO) {

			j := 0

			for j < len(item.DATOSINFO[z].DEPARTINFO) {

				i := 0

				for i < len(item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA) {

					p3 := item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i].CALIFICACION

					n := item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i].NOMBRE
					//d := item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i].DESCRIPCION
					c := item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i].CONTACTO

					n1 := item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i+1].NOMBRE

					s := strconv.Itoa(p3)

					g.addEdge(n, n1, "sig")
					g.addEdge(n, c, "contacto")
					g.addEdge(n, s, "calificacion")
					g.addEdge(n1, n, "ante")

					i = i + 1
				}

				j = j + 1

			}

			z = z + 1

		}

	}
	fmt.Println(g)

	path2, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path2, "-Tpng", "grafo.dot").Output()
	mode := int(0777)
	ioutil.WriteFile("grafo.png", cmd, os.FileMode(mode))

}


// }no usar esta seccion encacillada





func main() {
	router := mux.NewRouter().StrictSlash(true)

	//endpoints
	router.HandleFunc("/cargartienda", ListadoDeTiendasEndpoint).Methods("GET")

	router.HandleFunc("/cargarinventario", ListadoDeInventarioEndpoint).Methods("GET")

	router.HandleFunc("/cargarpedido", ListadoDePedidoEndpoint).Methods("GET")



	router.HandleFunc("/cargarusuarios", ListadoDeUsuariosEndpoint).Methods("GET")

	router.HandleFunc("/cargarusuarios/{Cuenta}", ObtenerTipoCuentaEndpoint).Methods("GET")

	router.HandleFunc("/nuevoUsuarios", CrearUsuarioEndPoint).Methods("POST")

	router.HandleFunc("/EliminarUsuario", EliminarUsuarioEndPoint).Methods("POST")

	router.HandleFunc("/InicioSesion", InicioSesionEndPoint).Methods("POST")


	router.HandleFunc("/subirGrafo", SubirGrafoEndPoint).Methods("POST")

	router.HandleFunc("/verDatosGrafo", DatosGrafoEndPoint).Methods("GET")





	router.HandleFunc("/cargartienda/indice/{Indice}", ObtenerIndiceTiendaEndpoint).Methods("GET")
	router.HandleFunc("/cargartienda/tipotienda/{Nombre}", ObtenerNombreTipoTiendaEndpoint).Methods("GET")
	router.HandleFunc("/cargartienda/nombre/{Nombre}", ObtenerNombreTiendaEndpoint).Methods("GET")
	router.HandleFunc("/cargartienda/calificacion/{Contacto}", ObtenerContactoTiendaEndpoint).Methods("GET")

	//---router.HandleFunc("/cargartienda/{Id nuevo}", CrearTiendaEndpoint).Methods("POST")

	//prueba post
	router.HandleFunc("/cargartienda/subir", CrearEndPoint).Methods("POST")

	router.HandleFunc("/cargartienda/{Nombre}", BorrarTiendaEndpoint).Methods("DELETE")

	router.HandleFunc("/TiendaEspecifica", TiendaEspecificaEndPoint).Methods("POST")

	router.HandleFunc("/Eliminar", EliminarRegistroEndPoint).Methods("POST")



	router.HandleFunc("/Inventario", InventarioEndPoint).Methods("POST")
	router.HandleFunc("/Pedidos", PedidosEndPoint).Methods("POST")




    // no usar esta seccion encacillada{

	router.HandleFunc("/Descargar", DescargarArchivoEndPoint)

	router.HandleFunc("/orden", Ordenamiento).Methods("GET")

	router.HandleFunc("/getArreglo", Grafica).Methods("GET")

	// }no usar esta seccion encacillada





	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}






// no usar esta seccion encacillada{

type edge struct {
	node  string
	label string
}
type graph struct {
	nodes map[string][]edge
}

func newGraph() *graph {
	return &graph{nodes: make(map[string][]edge)}
}

func (g *graph) addEdge(from, to, label string) {
	g.nodes[from] = append(g.nodes[from], edge{node: to, label: label})
}

func (g *graph) getEdges(node string) []edge {
	return g.nodes[node]
}

func (e *edge) String() string {
	return fmt.Sprintf("%v", e.node)
}

func (g *graph) String() string {
	out := `digraph finite_state_machine {
		rankdir=LR;
		size="8,5"
		node [shape = circle];`
	for k := range g.nodes {
		for _, v := range g.getEdges(k) {
			out += fmt.Sprintf("\t%s -> %s\t[ label = \"%s\" ];\n", k, v.node, v.label)
		}
	}
	out += "}"
	return out
}
// }no usar esta seccion encacillada