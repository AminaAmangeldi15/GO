package main

import (
	"database/sql"
	"strconv"

	// "html/template"
	"encoding/json"
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

type JsonResponse1 struct {
	Type    string `json:"type"`
	Data    []Item `json:"data1"`
	Message string `json:"message"`
}

type Filtering struct {
	From int `json:"item_f"`
	To int `json:"item_t"`
}

type Rating struct {
	Item_name string
	Rating    float64
}

func Db() *sql.DB {
	connStr := "user=postgres password=9792amina dbname=golang sslmode=disable"
	DB, err := sql.Open("postgres", connStr)
	CheckErr(err)
	// defer DB.Close()
	return DB
}

func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}


func Authorize(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("user_login")
	password := r.FormValue("user_password")

	var response = JsonResponse{}

	if login == "" || password == "" {
		response = JsonResponse{Type: "error", Message: "You are missing login or password parameter."}
	} else {
		db := Db()
		fmt.Println("Trying to find user with login: " + login)
		result, err := db.Exec("select * from users where login = $1 and password = $2", login, password)
		CheckErr(err)

		rows, err := db.Query("select * from users where login = $1 and password = $2", login, password)
		if err != nil {
			panic(err)
		}
		users := []User{}
		// items1 := []string{}
		for rows.Next(){
			i := User{}
			err := rows.Scan(&i.Name, &i.Surname, &i.Login, &i.Password)
			if err != nil{
				fmt.Println(err)
				continue
			}
			users  = append(users , i)
			// items1 = append(items1, fmt.Sprintf("item: ", i))
		}

	if len(users) == 0 {
		response = JsonResponse{Type: "error", Message: "The user is not found!"}
		} else {
			response = JsonResponse{Type: "success", Message: "You authorized!"}
		}
		result.LastInsertId()
	}
	json.NewEncoder(w).Encode(response)

}

func Register(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("user_name")
	surname := r.FormValue("user_surname")
	login := r.FormValue("user_login")
	password := r.FormValue("user_password")

	var response = JsonResponse{}

	if name == "" || surname == "" || login == "" || password == "" {
		response = JsonResponse{Type: "error", Message: "You are missing name or surname or login or password parameter."}
	} else {
		db := Db()

		fmt.Println("Inserting new user with name and surname: " + name + " " + surname + " and login: " + login)
		newRegister := User{Name: name, Surname: surname, Login: login, Password: password}
		result, err := db.Exec("insert into users (name, surname, login, password) values ($1, $2, $3, $4)", newRegister.Name, newRegister.Surname, newRegister.Login, newRegister.Password)
		CheckErr(err)
		result.LastInsertId()
		response = JsonResponse{Type: "success", Message: "The user has been inserted successfully!"}
	}
	json.NewEncoder(w).Encode(response)

}

func AddItem(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("item_name")
	price := r.FormValue("item_price")
	rating := r.FormValue("item_rating")
	var response = JsonResponse1{}

	p, err := strconv.ParseFloat(price, 64)
	CheckErr(err)
	rate, err := strconv.ParseFloat(rating, 64)
	CheckErr(err)
	if name == "" || price == "" || rating == "" {
		response = JsonResponse1{Type: "error", Message: "You are missing name or price or rating parameter."}
	} else {
		db := Db()

		fmt.Println("Inserting new item with name and price: " + name + " " + price)
		newRegister := Item{Name: name, Price: p, Rating: rate}
		result, err := db.Exec("insert into items (name, price, rating) values ($1, $2, $3)", newRegister.Name, newRegister.Price, newRegister.Rating)
		CheckErr(err)
		result.LastInsertId()
		response = JsonResponse1{Type: "success", Message: "The item has been inserted successfully!"}
	}
	json.NewEncoder(w).Encode(response)
}

func SearchItemByName(w http.ResponseWriter, r *http.Request){
	name := r.FormValue("item_name")
	var response = JsonResponse1{}
	if name == ""  {
		response = JsonResponse1{Type: "error", Message: "You are missing name parameter."}
	} else {
		db := Db()
		fmt.Println("Searching item by name : " + name)
		result, err := db.Exec("select * from items where name like $1", name)
		CheckErr(err)

		rows, err := db.Query("select * from items where name like $1", name)
		CheckErr(err)
		
		items := []Item{}
		// items1 := []string{}
		for rows.Next(){
			i := Item{}
			err := rows.Scan(&i.Name, &i.Price, &i.Rating)
			if err != nil{
				fmt.Println(err)
				continue
			}
			items  = append(items , i)
			// items1 = append(items1, fmt.Sprintf("item: ", i))
		}
		items1 := fmt.Sprintf("item: ", items)

		result.LastInsertId()
		if len(items) == 0 {
			response = JsonResponse1{Type: "error", Message: "The item is not found!"}
		} else {
			response = JsonResponse1{Type: "success", Message: items1}
		}
	}
	json.NewEncoder(w).Encode(response)
}

func FilterByRating(w http.ResponseWriter, r *http.Request){
	from := r.FormValue("item_f")
	to := r.FormValue("item_t")
	var response = JsonResponse1{}
	if from == "" || to ==""  {
		response = JsonResponse1{Type: "error", Message: "You are missing from or to parameter."}
	} else {
		db := Db()
		fmt.Println("Searching item by rating from   " + from + " to " + to)
		result, err := db.Exec("select * from items where rating between $1 and $2", from, to)
		CheckErr(err)

		rows, err := db.Query("select * from items where rating between $1 and $2", from, to)
		if err != nil {
			panic(err)
		}
		items := []Item{}
		// items1 := []string{}
		for rows.Next(){
			i := Item{}
			err := rows.Scan(&i.Name, &i.Price, &i.Rating)
			if err != nil{
				fmt.Println(err)
				continue
			}
			items  = append(items , i)
			// items1 = append(items1, fmt.Sprintf("item: ", i))
		}
		items1 := fmt.Sprintf("item: ", items)

		result.LastInsertId()
		if len(items) == 0 {
			response = JsonResponse1{Type: "error", Message: "The item is not found!"}
		} else {
			response = JsonResponse1{Type: "success", Message: items1}
		}
	}
	json.NewEncoder(w).Encode(response)
}
	
func FilterByPrice(w http.ResponseWriter, r *http.Request){
	from := r.FormValue("item_f")
	to := r.FormValue("item_t")
	var response = JsonResponse1{}
	if from == "" || to ==""  {
		response = JsonResponse1{Type: "error", Message: "You are missing from or to parameter."}
	} else {
		db := Db()
		fmt.Println("Searching item by price from   " + from + " to " + to)
		result, err := db.Exec("select * from items where price between $1 and $2", from, to)
		CheckErr(err)

		rows, err := db.Query("select * from items where price between $1 and $2", from, to)
		CheckErr(err)
		items := []Item{}
		// items1 := []string{}
		for rows.Next(){
			i := Item{}
			err := rows.Scan(&i.Name, &i.Price, &i.Rating)
			if err != nil{
				fmt.Println(err)
				continue
			}
			items  = append(items , i)
			// items1 = append(items1, fmt.Sprintf("item: ", i))
		}
		items1 := fmt.Sprintf("item: ", items)

		result.LastInsertId()
		if len(items) == 0 {
			response = JsonResponse1{Type: "error", Message: "The item is not found!"}
		} else {
			response = JsonResponse1{Type: "success", Message: items1}
		}
	}
	json.NewEncoder(w).Encode(response)
}

func Rate(w http.ResponseWriter, r *http.Request){
	name := r.FormValue("item_name")
	rating := r.FormValue("item_rating")
	var response = JsonResponse1{}
	if name == "" || rating ==""  {
		response = JsonResponse1{Type: "error", Message: "You are missing name or rating parameter."}
	} else {
		db := Db()
		fmt.Println("Rating item " + name + " with " + rating)
		result, err := db.Exec("insert into rating (item_name, rating) values ($1, $2)", name, rating)
		CheckErr(err)
		rows, err := db.Query("select * from rating where item_name like $1", name)
		if err != nil {
			panic(err)
		}
		
		ratings := []Rating{}
			
		for rows.Next(){
			r := Rating{}
			err := rows.Scan(&r.Item_name, &r.Rating)
			if err != nil{
				fmt.Println(err)
				continue
			}
			ratings = append(ratings, r)
		}

		var sum float64
		for _, r := range ratings{
			sum += r.Rating
		}
		rate := sum / float64(len(ratings ))
		row1, err := db.Query("update items set rating = $1 where name = $2", rate, name)
		CheckErr(err)
		row1.Close()
		result.LastInsertId()
	}
	json.NewEncoder(w).Encode(response)
}

// func Register(u m.User){

// 	result, err := db.Exec("insert into users (name, surname, login, password) values ($1, $2, $3, $4)", u.Name, u.Surname, u.Login, u.Password)
// 	if err != nil{
// 		panic(err)
// 	}
// 	// fmt.Println(result.LastInsertId())  
//     fmt.Println("You registered a new user")
// 	result.LastInsertId()
// 	func Register(w http.ResponseWriter, r *http.Request) {
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
// func Register(w http.ResponseWriter, r *http.Request) {
// 	name := r.FormValue("user_name")
// 	surname := r.FormValue("user_surname")
// 	login := r.FormValue("user_login")
// 	password := r.FormValue("user_password")

// 	var response = m.JsonResponse{}

// 	if name == "" || surname == "" || login == "" || password == "" {
// 		response = m.JsonResponse{Type: "error", Message: "You are missing name or surname or login or password parameter."}
// 	} else {
// 		// db := Db()

// 		fmt.Println("Inserting new user with name and surname: " + name + " " + surname + " and login: " + login)
// 		newRegister := m.User{Name: name, Surname: surname, Login: login, Password: password}
// 		result, err := db.Exec("insert into users (name, surname, login, password) values ($1, $2, $3, $4)", newRegister.Name, newRegister.Surname, newRegister.Login, newRegister.Password)
// 		if err != nil {
// 			panic(err)
// 		}
// 		result.LastInsertId()
// 		response = m.JsonResponse{Type: "success", Message: "The user has been inserted successfully!"}
// 	}
// 	json.NewEncoder(w).Encode(response)

// }


// func Authorize(a m.Authorization){
// 	result, err := db.Exec("select * from users where login = $1 and password = $2", a.Login, a.Password)
// 	if err != nil{
// 		panic(err)
// 	}
// 	fmt.Println("You autorized!")
// 	result.LastInsertId()
// }

// func AddItem(i m.Item){
// 	result,err := db.Exec("insert into items (name, price, rating) values ($1, $2, $3)", i.Name, i.Price, i.Rating)
//     if err != nil{
//         panic(err)
//     }
// 	fmt.Println("You added a new item")
// 	result.LastInsertId()
// }

// func SearchItemByName(name string){
// 	rows, err := db.Query("select * from items where name like $1", name)
//     if err != nil {
//         panic(err)
//     }
    
//     items := []m.Item{}
     
//     for rows.Next(){
//         i := m.Item{}
//         err := rows.Scan(&i.Name, &i.Price, &i.Rating)
//         if err != nil{
//             fmt.Println(err)
//             continue
//         }
//         items  = append(items , i)
//     }

// 	fmt.Println("Searching item is:")
//     for _, i := range items {
//         fmt.Println(i.Name, i.Price, i.Rating)
//     }
// }

// func FilterByPrice(from, to int){
// 	rows, err := db.Query("select * from items where price between $1 and $2", from, to)
//     if err != nil {
//         panic(err)
//     }
    
//     items := []m.Item{}
     
//     for rows.Next(){
//         i := m.Item{}
//         err := rows.Scan(&i.Name, &i.Price, &i.Rating)
//         if err != nil{
//             fmt.Println(err)
//             continue
//         }
//         items  = append(items , i)
//     }

// 	if len(items) == 0{
// 		fmt.Printf("Searching items, which price is from %s to %s do not exist...\n", strconv.Itoa(from), strconv.Itoa(to))
// 	} else {
// 		fmt.Printf("Searching items, which price is from %s to %s are:\n", strconv.Itoa(from), strconv.Itoa(to))
// 		for _, i := range items {
// 			fmt.Println(i.Name, i.Price, i.Rating)
// 		}
// 	}
// }

// func FilterByRating(from, to int){
// 	rows, err := db.Query("select * from items where rating between $1 and $2", from, to)
//     if err != nil {
//         panic(err)
//     }
    
//     items := []m.Item{}
     
//     for rows.Next(){
//         i := m.Item{}
//         err := rows.Scan(&i.Name, &i.Price, &i.Rating)
//         if err != nil{
//             fmt.Println(err)
//             continue
//         }
//         items  = append(items , i)
//     }

// 	if len(items) == 0{
// 		fmt.Printf("Searching items, which rating is from %s to %s do not exist...\n", strconv.Itoa(from), strconv.Itoa(to))
// 	} else {
// 		fmt.Printf("Searching items, which rating is from %s to %s are:\n", strconv.Itoa(from), strconv.Itoa(to))
// 		for _, i := range items {
// 			fmt.Println(i.Name, i.Price, i.Rating)
// 		}
// 	}
// }

// func Rate(r m.Rating){
// 	result,err := db.Exec("insert into rating (item_name, rating) values ($1, $2)", r.Item_name, r.Rating)
// 	if err != nil{
// 		panic(err)
// 	}
// 	result.LastInsertId()
// 	rows, err := db.Query("select * from rating where item_name like $1", r.Item_name)
// 	if err != nil {
// 		panic(err)
// 	}
	
// 	ratings := []m.Rating{}
		
// 	for rows.Next(){
// 		r := m.Rating{}
// 		err := rows.Scan(&r.Item_name, &r.Rating)
// 		if err != nil{
// 			fmt.Println(err)
// 			continue
// 		}
// 		ratings = append(ratings, r)
// 	}
// 	var sum float64
// 	for _, r := range ratings{
// 		sum += r.Rating
// 	}
// 	rate := sum / float64(len(ratings ))
// 	row1, err := db.Query("update items set rating = $1 where name = $2", rate, r.Item_name)
// 	if err != nil{
// 		panic(err)
// 	}
// 	row1.Close()
// }

func main() {

	http.HandleFunc("/register", Register)
	http.HandleFunc("/authorize", Authorize)
	http.HandleFunc("/additem", AddItem)
	http.HandleFunc("/byname", SearchItemByName)
	http.HandleFunc("/rating", FilterByRating)
	http.HandleFunc("/price", FilterByPrice)
	http.HandleFunc("/giverating", Rate)
	fmt.Println("Server: http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
