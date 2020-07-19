package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/Projects/CabService_RestApi/db"

	"github.com/gorilla/mux"
	"github.com/Projects/CabService_RestApi/models"
)

//GetCabs : get closest cabs
func GetCabs(w http.ResponseWriter, r *http.Request) {

	var nearbycabs []models.Cab

	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	strVar := params["id"]
	Customerid, _ := strconv.Atoi(strVar)
	var customer models.Customer
	var db1 *gorm.DB = db.Init()

	if err := db1.Where("customerid =? ", Customerid).Find(&customer).Error; err != nil {
		http.Error(w, "Customer has not registered!", http.StatusBadRequest)
		return
	}
	fmt.Println(customer.Startlatitude, customer.Startlongitude)

	if err := db1.Where("SQRT( ABS(presentlatitude - ? )+ABS(presentlongitude - ? )) < ? AND status = ?", customer.Startlatitude, customer.Startlongitude, 2, true).Find(&nearbycabs).Error; err != nil {
		http.Error(w, "There are no vehicles around", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(nearbycabs)

}

//BookCab : Customer can book his cab
func BookCab(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	strVar := params["id"]
	Customerid, _ := strconv.Atoi(strVar)
	var customer models.Customer
	var db1 *gorm.DB = db.Init()

	if err := db1.Where("customerid =? ", Customerid).Find(&customer).Error; err != nil {
		http.Error(w, "Customer has not registered!", http.StatusBadRequest)
		return
	}

	var cab models.Cab
	//If both start and end location same.
	if customer.Startlatitude == customer.Endlatitude && customer.Startlongitude == customer.Endlongitude {
		http.Error(w, "Start and End Location cannot be same.", http.StatusBadRequest)
		return
	}

	if err := db1.Where("SQRT( ABS(presentlatitude - ? )+ABS(presentlongitude - ? )) < ? AND status = ?", customer.Startlatitude, customer.Startlongitude, 2, true).First(&cab).Error; err != nil {
		http.Error(w, "No cabs around you,try again later.", http.StatusBadRequest)
		return
	}

	var customerhistory models.Customerhistory
	customerhistory = models.Customerhistory{Customerid: Customerid, Startlatitude: customer.Startlatitude, Startlongitude: customer.Startlatitude, Endlatitude: customer.Endlatitude, Endlongitude: customer.Endlongitude, Cabid: cab.Cabid}
	db1.Debug().Create(&customerhistory)

	db1.Model(&cab).Where("cabid=?", cab.Cabid).Update(map[string]interface{}{"customerid": Customerid, "bookingid": customerhistory.Bookingid, "status": false})

	json.NewEncoder(w).Encode(cab)

}

//RideComplete : Drive ends ride.
func RideComplete(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	strVar := params["id"]
	Cabid, _ := strconv.Atoi(strVar)
	var cab models.Cab
	var db1 *gorm.DB = db.Init()
	db1.Preload("cab").Where("cabid =? ", Cabid).Find(&cab)
	db1.Model(&cab).Where("cabid=?", cab.Cabid).Update(map[string]interface{}{"customerid": 0, "bookingid": 0, "status": true})
	json.NewEncoder(w).Encode(cab)

}

//GetRideHistory : Get Customer's ride history
func GetRideHistory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	strVar := params["id"]
	Customerid, _ := strconv.Atoi(strVar)
	var customer models.Customer
	var db1 *gorm.DB = db.Init()
	if err := db1.Where("customerid =? ", Customerid).Find(&customer).Error; err != nil {
		http.Error(w, "Customer has not registered!", http.StatusBadRequest)
		return
	}

	var customerhistory []models.Customerhistory

	if err := db1.Where("customerid= ? ", Customerid).Find(&customerhistory).Error; err != nil {
		http.Error(w, "Customer has not taken a ride yet!", http.StatusBadRequest)
		return

	}
	json.NewEncoder(w).Encode(customerhistory)
}
