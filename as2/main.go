package main

import (
	"database/sql"
	// "encoding/json"
	"fmt"
	// "strconv"

	// "log"
	"net/http"

	_ "github.com/lib/pq"
	// a "github.com/AminaAmangeldi15/Go/as"
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
	Name   string  `json:"item_name"`
	Price  float64 `json:"item_price"`
	Rating float64 `json:"item_rating"`
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

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "This is homepage!")
// 	fmt.Println("Endpoint hit: homePage")
// }

func Register(w http.ResponseWriter, r *http.Request) {
	db := Db()
	// var Name, Username, Password string
	if r.Method == http.MethodGet {
	//   temp, _ := template.ParseFiles("views/register.html")
	//   temp.Execute(w, nil)
	}else if r.Method == http.MethodPost {
	  // fmt.Println("lpad")
	  Name := r.Form.Get("user_name")
	  Surname := r.Form.Get("user_surname")
	  Username := r.Form.Get("user_login")
	  Password := r.Form.Get("user_password")
  
	  fmt.Println(Name)
	  rows, err := db.Query("insert into users (name, surname, login, password) values (?,?,?,?)", Name, Surname,  Username, Password)
	  if err != nil {
		panic(err)
		// fmt.Println("You have entered an incorrect login or password, please try again!")
	  }else {
		fmt.Println("Okayy!")
	}
	rows.Close()
  
	}
}
// func createPost(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var post User
// 	_ = json.NewDecoder(r.Body).Decode(&post)
// 	post.Name = strconv.Itoa(rand.Intn(1000000))
// 	posts = append(posts, post)
// 	json.NewEncoder(w).Encode(&post)
// }

// func SearchItemByName(w http.ResponseWriter, r *http.Request){
// 	params := mux.Vars(r)
// 	rows, err := Db().Query("select * from items")
//     if err != nil {
//         panic(err)
//     }

//     items := []Item{}

//     for rows.Next(){
//         i := Item{}
//         err := rows.Scan(&i.Name, &i.Price, &i.Rating)
//         if err != nil{
//             fmt.Println(err)
//             continue
//         }
//         items  = append(items , i)
//     }

// 	w.Header().Set("Content-Type", "application/json")
//   	for _, item := range items {
// 		s := fmt.Sprintf("%f", item.Price) 
// 		if s == params["item_price"] {
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
//     }
//   json.NewEncoder(w).Encode(&Item{})
// }
// 	json.NewEncoder(w).Encode(items)
// 	fmt.Println("Endpoint hit: getAll")
// fmt.Println("Searching item is:")
// for _, i := range items {
//     fmt.Println(i.Name, i.Price, i.Rating)
// }
// }
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

// func Register(w http.ResponseWriter, r *http.Request) {
// 	name := r.FormValue("user_name")
// 	surname := r.FormValue("user_surname")
// 	login := r.FormValue("user_login")
// 	password := r.FormValue("user_password")

// 	var response = JsonResponse{}

// 	if name == "" || surname == "" || login == "" || password == "" {
// 		response = JsonResponse{Type: "error", Message: "You are missing name or surname or login or password parameter."}
// 	} else {
// 		db := Db()

// 		fmt.Println("Inserting new user with name and surname: " + name + " " + surname + " and login: " + login)
// 		newRegister := User{Name: name, Surname: surname, Login: login, Password: password}
// 		result, err := db.Exec("insert into users (name, surname, login, password) values ($1, $2, $3, $4)", newRegister.Name, newRegister.Surname, newRegister.Login, newRegister.Password)
// 		if err != nil {
// 			panic(err)
// 		}
// 		result.LastInsertId()
// 		response = JsonResponse{Type: "success", Message: "The user has been inserted successfully!"}
// 	}
// 	json.NewEncoder(w).Encode(response)

// }

func main() {
	// http.HandleFunc("/", homePage)
	// http.HandleFunc("/users", Register)

	// http.HandleFunc("/users", Register)
	// fmt.Println("Server at 8181")
	// err := http.ListenAndServe(":8181", nil)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// router := mux.NewRouter()
	// posts = append(posts, Post{ID: "1", Title: "My first post", Body:      "This is the content of my first post"})
	// router.HandleFunc("/posts", getPosts).Methods("GET")
	// router.HandleFunc("/posts/{item_name}", SearchItemByName).Methods("GET")
	// router.HandleFunc("/posts", Register).Methods("POST")
	// http.ListenAndServe(":8000", router)
	http.HandleFunc("/register", Register)

	fmt.Println("Server: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
	// router := mux.NewRouter()
	// router.HandleFunc("/users", Register).Methods("POST")
	// fmt.Println("Server at 8080")
	// log.Fatal(http.ListenAndServe(":8000", router))
}
