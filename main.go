package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Car struct {
	ID         int
	Brand      string
	Model      string
	HorsePower int
}

func carPosterHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		//w.WriteHeader(405)
		fmt.Fprintln(w, "Nothing to GET here")

	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatalln(err)
		}

		var car Car
		json.Unmarshal([]byte(body), &car)
		car = generateID(car)
		log.Printf("Car added successfully: %+v\n", car)
		fmt.Fprintf(w, "Car added successfully: %+v\n", car)
		storeInDatabase(car)

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func generateID(car Car) Car {
	rand.Seed(time.Now().UnixNano())
	ID := rand.Intn(1000000)
	car.ID = ID
	return car
}

func storeInDatabase(car Car) {
	db, err := sql.Open("mysql", "remoteuser:a@tcp(mariadb1:3306)/entrust")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var msg string = "INSERT INTO car (id, brand, model, horsePower) VALUES (" + strconv.Itoa(car.ID) + ", ' " + car.Brand + "', '" + car.Model + "', " + strconv.Itoa(car.HorsePower) + ");"

	insert, err := db.Query(msg)
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
}

func main() {
	fmt.Println("POST Server Running...	Hola2")
	http.HandleFunc("/service/v1/cars", carPosterHandler)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
