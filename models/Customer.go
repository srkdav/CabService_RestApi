package models

// package main

// Customer is
type Customer struct {
	Customerid     int    `json:"customerid" gorm:"primary_key;AUTO_INCREMENT:true"`
	Customername   string `json:"customername"`
	Startlongitude int    `json:"startlongitude"`
	Startlatitude  int    `json:"startlatitude"`
	Endlongitude   int    `json:"endlongitude"`
	Endlatitude    int    `json:"endlatitude"`
	Cab            *Cab   `json:"cab" gorm:"foreignkey:Customerid"`
}
