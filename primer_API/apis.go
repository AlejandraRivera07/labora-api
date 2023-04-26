package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var items []Item = []Item{
	{
		ID:   1,
		Name: "Camila",
	},
	{
		ID:   2,
		Name: "Paula",
	},
	{
		ID:   3,
		Name: "Alejandra",
	},
	{
		ID:   4,
		Name: "Andres",
	},
	{
		ID:   5,
		Name: "Luis",
	},
	{
		ID:   6,
		Name: "Camilo",
	},
	{
		ID:   7,
		Name: "Luisa",
	},
	{
		ID:   8,
		Name: "Juan",
	},
	{
		ID:   9,
		Name: "Liz",
	},
	{
		ID:   10,
		Name: "Carmen",
	},
}

func obtenerItems(response http.ResponseWriter, request *http.Request) {
	// creamos una estructura para representar nuestros datos

	jsonData, _ := json.Marshal(items)
	response.WriteHeader(http.StatusOK)
	response.Write([]byte(jsonData))
}

func buscarID(response http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	idStr := vars["id"]

	// Convertir el ID de string a int
	id, _ := strconv.Atoi(idStr)
	//id := getQueryParam(request, "id")

	// Buscar el elemento con el id que corresponde
	//var itemByID *Item
	var itemName string
	for i := 0; i < len(items); i++ {
		if items[i].ID == id {
			itemName = items[i].Name
			break
		}
	}
	if itemName == "" {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte("No se encontró el elemento con ID " + idStr))
		return

	}

	jsonData, err := json.Marshal(itemName)
	if err != nil {
		// Manejar el error de la conversión a JSON
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("Error al convertir a JSON"))
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write([]byte(jsonData))

}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/items", obtenerItems).Methods("GET")
	router.HandleFunc("/items/{id}", buscarID).Methods("GET")
	direccion := ":8080"
	servidor := &http.Server{
		Handler: router,
		Addr:    direccion,
		// Timeouts para evitar que el servidor se quede "colgado" por siempre
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Printf("Escuchando en %s. Presiona CTRL + C para salir", direccion)
	log.Fatal(servidor.ListenAndServe())
	// listen to port
	http.ListenAndServe(":8080", nil)
}
