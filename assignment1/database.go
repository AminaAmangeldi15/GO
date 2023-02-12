package assignment1

import (
	"fmt"
	"strings"
)

// import "github.com/aminaamangeldi15/go/assignment1"
// import (
// 	"./assignment1"
// )
// import "C:\Users\amina\KBTU\Go\assignment1\registration.go"

// type Database struct {
// 	Logins []Registration
// }

type Database struct {
	Logins []Registration
	Items []Item
}

func (d *Database) AddItem(i Item) *Item{
	item := &Item{i.Name, i.Price, i.Rating}
	d.Items = append(d.Items, Item{i.Name, i.Price, i.Rating})
	return item
}

func (d *Database) SearchItem(name string){
	for i := 0; i < len(d.Items); i++{
		if strings.Contains(strings.ToUpper(d.Items[i].Name), strings.ToUpper(name)){
			fmt.Println(d.Items[i])
			return
		} else {
			fmt.Printf("Did not find an item like %s", name)
		}
	}
}

func (d *Database) FilteringItems(price1, price2, rating1, rating2 float64){
	for i := 0; i < len(d.Items); i++{
		if d.Items[i].Price >= price1 && d.Items[i].Price <= price2 && d.Items[i].Rating >= rating1 && d.Items[i].Rating <= rating2{
			fmt.Println("We found item that you searched!")
			fmt.Println(d.Items[i])
			return
		} else {
			fmt.Println("No such item with these price and rating!!!")
		}
	}
}
// func (d *Database) Register(r Registration) *Registration {
// 	for i := 0; i < len(d.Logins); i++ {
// 		if d.Logins[i].Login == r.Login && d.Logins[i].Password == r.Password {
// 			// return fmt.Errorf("username %s %s already exists", r.Name, r.Surname)
// 			fmt.Println("User %s %s already exists", r.Name, r.Surname)
// 		}
// 	}
// 	reg := &Registration{r.Name, r.Surname, r.Age, r.Login, r.Password}
// 	d.Logins = append(d.Logins, Registration{r.Name, r.Surname, r.Age, r.Login, r.Password})
// 	fmt.Println("You registred!")
// 	return reg
// }

func (d *Database) Register(r Registration) *Registration {
	for i := 0; i < len(d.Logins); i++ {
		if d.Logins[i].Login == r.Login && d.Logins[i].Password == r.Password {
			// return fmt.Errorf("username %s %s already exists", r.Name, r.Surname)
			fmt.Printf("User %s %s already exists!\n", r.Name, r.Surname)
			return nil
		}
	}
	reg := &Registration{r.Name, r.Surname, r.Age, r.Login, r.Password}
	d.Logins = append(d.Logins, Registration{r.Name, r.Surname, r.Age, r.Login, r.Password})
	fmt.Println("You registred!")
	return reg
}

func Rate(d Database, a Authorization, itemName string, rating float64){
	if(a.SignIn(d) == "You entered system!"){
		fmt.Println("You entered system!")
		for i := 0; i < len(d.Items); i++{
			if d.Items[i].Name == itemName{
				d.Items[i].Rating = rating
				fmt.Println("You successfully rated an item!")
			} else {
				fmt.Println("No such item!!!")
			}
		}
	} else {
		fmt.Println("Wrong login or password!!!")
	}
}
