# Implementation of String Matching and Regular Expressions in DNA Pattern Matching

## Table of Contents
- [General Information](#general-information)
- [Requirement](#requirement)
- [How To Run](#how-to-run)
- [Screenshots](#screenshots)
- [Authors](#authors)

## General Information
This program is a website application program that using Golang for backend and VUE for frontend development. This website application is a DNA Pattern Matching application by utilizing string matching algorithms, Knuth-Morris-Pratt and Boyer Moore algorithm, as well as regular expressions.
<br>
This application is an interactive application that can detect whether a patient has a certain genetic disease and then store the test results in a database. This application can also add a list of genetic diseases in the database. In addition, this application can display all DNA tests based on the entered query

## Requirement
* [Node 14](https://nodejs.org/en/) or above
* [Vue 3](https://vuejs.org/)
* [Go 1.18](https://go.dev/)

## How To Run
* Clone this repository and install node modules:
  ```
  $ git clone https://github.com/adellinekania/Tubes3_dnAAAAA.git
  $ cd src/client
  $ npm install 
  ```
* Build the backend program:
  ```
  working directory: src/api
  $ go build
  ```
### Development Server
* Run golang backend server
  ```
  working directory: src/api (run executable result from build)
  $ .\tubes_3_dnAAAAA
  ```
* Run vue development server
  ```
  working directory: src/client
  $ npm run dev
  ```
  The Vue js application will be served from `localhost:3000` and the Go Backend will be served from `localhost:5000`.

  The dual dev-server setup allows you to take advantage of webpack's development server with hot module replacement.

  Proxy config in `vite.config.js` is used to route the requests back to go-mux Api on port 5000.

  If you would rather run a single dev server, you can run Go Mux server only on `:5000`, but you have to build the Vue app first and the page will not reload on     changes.
  ```
  working directory: src/client
  $ npm run build

  working directory: src/api
  $ .\tubes_3_dnAAAAA
  ```
  If you want to use the already deployed version, you can visit the following [link](https://dnaaaaa.jovaandreas.me/).
  
## Screenshots
<b>Landing Page</b>
![image](https://user-images.githubusercontent.com/64909665/165921505-68470bb0-18c5-4f75-947e-ff952c064aa5.png)
<b>Tambah Penyakit</b>
![image](https://user-images.githubusercontent.com/64909665/165921593-86a4d0b5-3ed9-472c-bd8f-ee31781e803f.png)
<b>Tes DNA</b>
![image](https://user-images.githubusercontent.com/64909665/165921641-dfb0789e-9305-4050-b6cd-b0e09f972129.png)
<b>Riwayat</b>
![image](https://user-images.githubusercontent.com/64909665/165921763-0e136042-0c7c-4195-a742-3a7ab20090bc.png)

## Authors

<b>dnAAAAA</b>
| NIM       | Name                      |
| --------- | --------------------------|
| 13520024  | Hilya Fadhilah Imania     |
| 13520072  | Jova Andreas Riski Sirait |
| 13520084  | Adelline Kania Setiyawan  |

TUGAS BESAR III - Algorithm Strategies
<br>
Bandung Institute of Technology
