package models

//Customerhistory struct is
type Customerhistory struct {
	Bookingid      int `json:"bookingid" gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Customerid     int `json:"customerid"`
	Cabid          int `json:"cabid"`
	Startlongitude int `json:"startlongitude"`
	Startlatitude  int `json:"startlatitude"`
	Endlongitude   int `json:"endlongitude"`
	Endlatitude    int `json:"endlatitude"`
}
