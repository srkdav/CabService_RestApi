

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
4.Change the MySQL properties :

```
db, err = gorm.Open("mysql", "<username>:<password>@tcp(localhost:<port number>)/<databse_name>")
	
```

5. Build && Run the EXE file 
```sh
  go build && ./CabService_RestApi
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


