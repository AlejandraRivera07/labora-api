package main

import (
	"conectar_db_api/config"
	"conectar_db_api/controllers"
	"log"

	"github.com/gorilla/mux"
)

func main() {
	//services.UpDb() //Esta llamada será para levantar la DB, la veremos mas adelante

	//Instancia del Router con Mux
	router := mux.NewRouter()

	//Endpoints
	router.HandleFunc("/items", controllers.GetItems).Methods("GET")

	router.HandleFunc("/items/{id}", controllers.GetItemById).Methods("GET")

	router.HandleFunc("/items/search/customer", controllers.getItemByName).Methods("GET")

	//services.Db.PingOrDie() //Checkeador de la DB, mas adelante veremos esto

	//levantamos el servidor
	port := ":9000"
	if err := config.StartServer(port, router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
