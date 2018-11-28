package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/handler"
)

// import "handler"

func main() {
	router := mux.NewRouter()

	log.Println("Server started on: http://localhost:8080")

	router.HandleFunc("/Registration", handler.Registration)

	router.HandleFunc("/Login", handler.Login)                             //Login
	router.HandleFunc("/ShowAddress/{userid}", handler.ShowAddress)        //Show Addresses
	router.HandleFunc("/AddAddress", handler.AddAddress)                   // Add Address
	router.HandleFunc("/UpdateAddress/{addressid}", handler.UpdateAddress) //Update Address
	router.HandleFunc("/DeleteAddress/{addressid}", handler.DeleteAddress) //Delete Address
	http.ListenAndServe(":8080", router)
}
