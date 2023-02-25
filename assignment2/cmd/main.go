package main

import (
	f "github.com/AminaAmangeldi15/Go/assignment2/pkg"
)




func main(){
	// result, err := db.Exec("insert into users (name, surname, login, password) values ('Amina', 'Amangeldi', 'login', 'passwprd')")
	// if err != nil{
	// 	panic(err)
	// }
	// u := User{"a", "a", "l1", "p1"}
	// // register(u)
	// Register(u)
	// Authorize(Authorization{"ff", "d"})
	// AddItem(Item{"Iphone 12", 15000, 0})
	// SearchItemByName("Iphone 12")
	// Rate(Rating{"Iphone 12", 4})
	// Rate(Rating{"Iphone 12", 3})
	// Rate(Rating{"Iphone 12", 2})
	f.SearchItemByName("Iphone 12")
	// FilterByPrice(10000, 200000)
	f.FilterByRating(2, 4)

}