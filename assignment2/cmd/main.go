package main

import (
	f "github.com/AminaAmangeldi15/Go/assignment2/pkg"
	m "github.com/AminaAmangeldi15/Go/assignment2/models"
)

func main(){
	// result, err := db.Exec("insert into users (name, surname, login, password) values ('Amina', 'Amangeldi', 'login', 'passwprd')")
	// if err != nil{
	// 	panic(err)
	// }
	// u := m.User{"name", "surname", "lap", "djk"}
	// f.Register(u)
	// f.AddItem(m.Item{"Samsung S22", 450000, 0})
	f.Rate(m.Rating{"Samsung S22", 5})
	f.SearchItemByName("Samsung S22")
	// Authorize(Authorization{"ff", "d"})
	// AddItem(Item{"Iphone 12", 15000, 0})
	// SearchItemByName("Iphone 12")
	// Rate(Rating{"Iphone 12", 4})
	// Rate(Rating{"Iphone 12", 3})
	// Rate(Rating{"Iphone 12", 2})
	// f.SearchItemByName("Iphone 12")
	// FilterByPrice(10000, 200000)
	// f.FilterByRating(2, 4)

}