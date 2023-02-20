package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Tutor struct {
	Name    string   `json: "name"`
	Surname string   `json: "surname"`
	Age     int      `json: "age"`
	Courses []Course `json: "courses"`
}

type Course struct {
	Title   string `json: "title"`
	Credits int    `json: "credits"`
	Tutor   Tutor  `json: "tutor"`
}

var Tutors []Tutor
var Courses []Course

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is a homepage")
	fmt.Println("endpoint hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/tutors", getAllTutors)
	fmt.Println("Server is listening...")
	err := http.ListenAndServe("8181", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func getAllTutors(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Tutors)
	fmt.Println("Endpoint hit: getAllTutors")
}

func main() {
	Courses = append(Courses, Course{"Math", 6, Tutor{}})
	t := Tutor{"Amina", "Amangeldi", 18, Courses}
	Tutors = append(Tutors, t)
	handleRequests()
}
