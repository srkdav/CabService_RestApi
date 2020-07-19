package models

// Cab Struct is :
type Cab struct {
	Cabid            int    `json:"cabid" gorm:"primary_key ;AUTO_INCREMENT"`
	Drivername       string `json:"drivername"`
	Presentlatitude  int    `json:"presentlatitude"`
	Presentlongitude int    `json:"presentlongitude"`
	Status           bool   `json:"status" gorm:"default:true"`
	Bookingid        int    `json:"bookingid"`
	Customerid       int    `json:"-"`
}
