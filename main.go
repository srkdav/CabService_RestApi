package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/Projects/CabService_RestApi/controller"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/getCabs/{id}", controller.GetCabs).Methods("POST")
	r.HandleFunc("/bookCab/{id}", controller.BookCab).Methods("POST")
	r.HandleFunc("/rideComplete/{id}", controller.RideComplete).Methods("POST")
	r.HandleFunc("/getHistory/{id}", controller.GetRideHistory).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
