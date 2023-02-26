package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// _ "github.com/lib/pq"
	a "github.com/AminaAmangeldi15/Go/as2"
	// "github.com/gorilla/mux"
	// "github.com/AminaAmangeldi15/Go/assignment2/pkg"
	// "github.com/AminaAmangeldi15/Go/assignment2/models"
)

type User struct {
	Name     string `json:"user_name"`
	Surname  string `json:"user_surname"`
	Login    string `json:"user_login"`
	Password string `json:"user_password"`
}

type JsonResponse struct {
	Type    string `json:"type"`
	Data    []User `json:"data"`
	Message string `json:"message"`
}

type Authorization struct {
	Login    string
	Password string
}

type Item struct {
	Name   string
	Price  float64
	Rating float64
}

type Rating struct {
	Item_name string
	Rating    float64
}

func Db() *sql.DB {
	connStr := "user=postgres password=9792amina dbname=golang sslmode=disable"
	DB, err := sql.Open("postgres", connStr)
	checkErr(err)
	// defer DB.Close()
	return DB
}

// func Register(u User){

//		result, err := db.Exec("insert into users (name, surname, login, password) values ($1, $2, $3, $4)", u.Name, u.Surname, u.Login, u.Password)
//		if err != nil{
//			panic(err)
//		}
//		// fmt.Println(result.LastInsertId())
//	    fmt.Println("You registered a new user")
//		result.LastInsertId()
//	}
func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	Name := r.FormValue("user_name")
	Surname := r.FormValue("user_surname")
	Login := r.FormValue("user_login")
	Password := r.FormValue("user_password")

	var response = JsonResponse{}

	if Name == "" || Surname == "" || Login == "" || Password == "" {
		response = JsonResponse{Type: "error", Message: "You are missing name or surname or login or password parameter."}
	} else {
		db := Db()

		printMessage("Inserting user into DB")

		fmt.Println("Inserting new user with name: " + Name + " and surname: " + Surname)

		var lastInsertID string
		err := db.QueryRow("insert into users (name, surname, login, password) values ($1, $2, $3, $4) returning *", Name, Surname, Login, Password).Scan(&lastInsertID)

		// check errors
		checkErr(err)

		response = JsonResponse{Type: "success", Message: "The user has been inserted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", Register).Methods("POST")

	fmt.Println("Server at 8181")
	log.Fatal(http.ListenAndServe(":8181", router))
}
