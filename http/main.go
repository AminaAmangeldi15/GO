package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	// "github.com/gorilla/mux"
)

type Article struct {
	Title string `json: "title"`
	Desc string `json: "desc"`
	Content string `json: "content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is homepage!")
	fmt.Println("Endpoint hit: homePage")
}

func getAll(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(Articles)
	fmt.Println("Endpoint hit: getAll")
}

func handleRequest(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/art", getAll)
	fmt.Println("Listenning")
	err := http.ListenAndServe(":8181", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
// func handleRequests() {
	// myRouter := mux.NewRouter().StrictSlash(true)
	// myRouter.HandleFunc("/", homePage)
	// myRouter.HandleFunc("/articles", getAll)
	// fmt.Println("Listenning...")
	// err := http.ListenAndServe(":8181", myRouter)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
// }

func main(){
	Articles = []Article{
		{"KBTU", "DESC KBTU", "CONTENT"},
		{"SDU", "DESC SDU", "CONTENT"},
	}
	handleRequest()
}