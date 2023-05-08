package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	//"strings"
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

func createItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	err := json.NewDecoder(r.Body).Decode(&item)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	items = append(items, item)

	w.Write([]byte("Item creado correctamente"))
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Convertir el ID de string a int
	id, _ := strconv.Atoi(idStr)
	var itemUpdate Item
	err := json.NewDecoder(r.Body).Decode(&itemUpdate)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, item := range items {
		if item.ID == id {
			items[i] = itemUpdate
			w.Write([]byte("Item actualizado correctamente"))
			return
		}
	}

	w.Write([]byte("No se pudo actualizar el item"))
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	// TODO Función para eliminar un elemento
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Convertir el ID de string a int
	id, _ := strconv.Atoi(idStr)
	for i, item := range items {
		if item.ID == id {
			nuevoSlice := make([]Item, len(items)-1)

			nuevoSlice = append(items[:i], items[i+1:]...)

			items = nuevoSlice
			w.Write([]byte("Item eliminado correctamente"))
			return
		}
	}

	w.Write([]byte("Item no pudo ser eliminado"))
}

func getItemByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	variables := r.URL.Query()
	name := variables.Get("name")

	for _, item := range items {
		if strings.ToLower(item.Name) == strings.ToLower(name) {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Item{})
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/items", obtenerItems).Methods("GET")
	router.HandleFunc("/items/{id}", buscarID).Methods("GET")
	router.HandleFunc("/items", createItem).Methods("POST")
	router.HandleFunc("/items/{id}", updateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")
	router.HandleFunc("/items/get-by-name/", getItemByName).Methods("GET")
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
	//http.ListenAndServe(":8080", nil)
}
