package models

type User struct {
	Name     string `json:"user_name"`
	Surname  string `json:"user_surname"`
	Login    string `json:"user_login"`
	Password string `json:"user_password"`
}

type Authorization struct {
	Login    string `json:"au_login"`
	Password string `json:"au_password"`
}

type Item struct {
	Name   string  `json:"item_name"`
	Price  float64 `json:"item_price"`
	Rating float64 `json:"item_rating"`
}

type Rating struct {
	Item_name string  `json:"r_name"`
	Rating    float64 `json:"r_rating"`
}
