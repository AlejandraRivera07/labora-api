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
