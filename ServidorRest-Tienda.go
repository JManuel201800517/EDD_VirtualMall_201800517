package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
}

var tiendaespecifica []Especifica

type Especifica struct {
	DepartamentoEspecifico string `json:"Departamento,omitempty"`
	NombreEspecifico       string `json:"Nombre,omitempty"`
	CalificacionEspecifico int    `json:"Calificacion,omitempty"`
}

func ListadoDeTiendasEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(cargartienda)

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

func main() {
	router := mux.NewRouter()

	//enspoints
	router.HandleFunc("/cargartienda", ListadoDeTiendasEndpoint).Methods("GET")

	router.HandleFunc("/cargartienda/indice/{Indice}", ObtenerIndiceTiendaEndpoint).Methods("GET")
	router.HandleFunc("/cargartienda/tipotienda/{Nombre}", ObtenerNombreTipoTiendaEndpoint).Methods("GET")
	router.HandleFunc("/cargartienda/nombre/{Nombre}", ObtenerNombreTiendaEndpoint).Methods("GET")
	router.HandleFunc("/cargartienda/calificacion/{Contacto}", ObtenerContactoTiendaEndpoint).Methods("GET")

	//---router.HandleFunc("/cargartienda/{Id nuevo}", CrearTiendaEndpoint).Methods("POST")

	//prueba post
	router.HandleFunc("/cargartienda/subir", CrearEndPoint).Methods("POST")

	router.HandleFunc("/cargartienda/{Nombre}", BorrarTiendaEndpoint).Methods("DELETE")

	router.HandleFunc("/TiendaEspecifica", TiendaEspecificaEndPoint).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", router))
}
