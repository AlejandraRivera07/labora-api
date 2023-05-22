package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"conectar_db_api/services"

	"github.com/gorilla/mux"
)

type ItemHandler struct {
	ItemService services.DbConnection
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	items, err := services.GetItems()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Internal Server Error")
		return
	}
	json.NewEncoder(w).Encode(items)

}

func GetItemById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Fatal(err)
	}

	item, err := services.GetItemById(id)

	if item == nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Objeto con id %d no encontrado", id)))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

func getItemByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	variables := r.URL.Query()
	customer, _ := variables["customer"]

	item, err := services.SearchItemByName(strings.ToLower(customer[0]))

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Internal Server Error")
		return
	}

	json.NewEncoder(w).Encode(item)
}

func EditItemByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Convertir el ID de string a int
	id, _ := strconv.Atoi(idStr)
	var itemUpdate models.Item
	_ = json.NewDecoder(r.Body).Decode(&itemUpdate)
	log.Printf("Item a actualizar: %+v", itemUpdate)
	itemUpdated, err := services.UpdateItemByID(id, itemUpdate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("Item a actualizado: %+v", itemUpdated)
	jsonData, err := json.Marshal(itemUpdated)
	if err != nil {
		log.Printf("Error %v", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonData))

}

func DeletItemByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Convertir el ID de string a int
	id, _ := strconv.Atoi(idStr)
	var itemDelet models.Item
	_ = json.NewDecoder(r.Body).Decode(&itemDelet)
	log.Printf("Item a borrar: %+v", itemDelet)
	itemDeleted, err := services.DeletItemByID(id, itemDelet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("Item que se borro: %+v", itemDeleted)
	jsonData, err := json.Marshal(itemDeleted)
	if err != nil {
		log.Printf("Error %v", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonData))
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var NewItem models.Item
	err := json.NewDecoder(r.Body).Decode(&NewItem)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Ocurrió un error al intentar crear el registro"))
		return
	}
	NewItem, err = services.CreateItem(NewItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(NewItem)
}

