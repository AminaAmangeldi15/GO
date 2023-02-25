package models

type User struct {
	Name string
	Surname string
	Login string
	Password string
}

type Authorization struct {
	Login    string
	Password string
}

type Item struct {
	Name       string
	Price      float64
	Rating     float64
}

type Rating struct{
	Item_name string
	Rating float64
}