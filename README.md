

<!-- PROJECT SHIELDS -->
<!-- PROJECT LOGO -->
<br />
<p align="center">
 
  <h3 align="center">Cab Service Rest API</h3>

  <p align="center">
   This application mimics the functioning of Uber/Ola. Customers can find closeby cabs,book the cloest available cab and also get the history of their rides.  
    <br />
  </p>
</p>



<!-- TABLE OF CONTENTS -->
## Table of Contents

* [About the Project](#about-the-project)
  * [Built With](#built-with)
* [Getting Started](#getting-started)
  * [Prerequisites](#prerequisites)
  * [Installation](#installation)
  * [After Running the Application](#after-running-the-application)
* [Contact](#contact)



<!-- ABOUT THE PROJECT -->
## About The Project

1. Customers can find and book a cab.
2. Drivers can also update their status.
3. Customers can find their ride history.

### Built With

* GoLang 1.14.5
* MySQL 


<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

```sh
* GoLang 1.14.5
* MySQL 
* PostMan
* VSCode
```

### Installation
 
1. GOPATH by default is C:/Users/<username>/go

2. Create a directory in the GOPATH :

```
mkdir -p $GOPATH/src/github.com/Projects
```
3. Clone the repo in VSCode in the "Projects" folder
```sh
git clone https://github.com/srkdav/CabService_RestApi.git
```
4.Change the MySQL properties under db/db.go :
```
create database cabrestapi;
```
```
db, err = gorm.Open("mysql", "<username>:<password>@tcp(localhost:<port number>)/cabrestapi")
	
```

5. Build && Run the EXE file 
```sh
  go build && ./CabService_RestApi
```

6. After running the application, Tables will be created under the "cabrestapi" Database.

7. Populate Customer data , Cab data :

Customer Data :
```
insert into customers values ( 1,"steve",0,0,2,2);
insert into customers values ( 2,"tony",2,2,4,4);
insert into customers values ( 3,"bruce",4,5,7,8);
insert into customers values ( 4,"bruce",4,4,4,4);
insert into customers values ( 5,"clark",10,10,15,16);
```

Cabs Data:
```
insert into cabs values ( 100,"lewis",1,1,true,0,0);
insert into cabs values ( 200,"dani",2,2,true,0,0);
insert into cabs values ( 300,"lewis",5,5,true,0,0);
insert into cabs values ( 400,"lewis",1,1,true,0,0);
insert into cabs values ( 500,"sainz",9,10,true,0,0);
```

### After Running the application

1. Customers can find the cabs around them : 
 ```
 http://localhost:8000/getCabs/{CustomerId} 
```
2. Customers can book an availble cab closest to them : 
```
http://localhost:8000/bookCab/{CustomerId}
```
3. Customers can also find their Ride history :
```
http://localhost:8000/getHistory/{CustomerId}
```
<!-- CONTACT -->
## Contact

Karthick Sabhapathy - www.linkedin.com/in/karthick-sabhapathy - email


