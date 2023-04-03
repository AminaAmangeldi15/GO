package pkg

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

func Routes(){
	router := mux.NewRouter()
	router.HandleFunc("/getById", GetBookById).Methods("GET")
	router.HandleFunc("/getAllBooks", GetAll).Methods("GET")
	router.HandleFunc("/updateById", UpdateById).Methods("PUT")
	router.HandleFunc("/deleteById", DeleteById).Methods("DELETE")
	router.HandleFunc("/searchByTitle", SearchByTitle).Methods("GET")
	router.HandleFunc("/addBook", AddBook).Methods("POST")
	router.HandleFunc("/getBooksByAsc", SortOrderByAsc).Methods("GET")
	router.HandleFunc("/getBooksByDesc", SortOrderByDesc).Methods("GET")
	fmt.Println("Server at 8080")
    errr := http.ListenAndServe(":8080", router)
	if errr != nil {
        fmt.Println(errr)
    }
}