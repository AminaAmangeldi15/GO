package pkg

import (
	"fmt"
	m "m/models"
	"encoding/json"
	"net/http"
	"gorm.io/gorm"
	"strconv"
)

var db *gorm.DB = DB()

func GetBookById(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	// db := DB()
	var response = m.JsonResponse{}
	book := m.Book{}
	// db.First(&book, id)
	db.Where("id = ? and id != '' and deleted_at is null", id).Find(&book)
	// db.Where("id=? and id != '' ", id).Find(&book)
	// fmt.Println(book)
	if id == "" {
        response = m.JsonResponse{Type: "error", Message: "You are missing id parameter."}
    } else if book.Id != ""{
		response = m.JsonResponse{Type: "success", Message: "Products: " , Data: book}
	} else {
		response = m.JsonResponse{Type: "error", Message: "Not found!"}
	}
	json.NewEncoder(w).Encode(response)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	// db := DB()
	var response = m.JsonResponse{}
	book := []m.Book{}
	db.Find(&book)
	// fmt.Println(book)
	response = m.JsonResponse{Type: "success", Message: "Products:" , Datas: book}
	json.NewEncoder(w).Encode(response)
}

func UpdateById(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	// db := DB()
	book := m.Book{}
	db.Where("id = ? and id != '' ", id).Find(&book)
	book.Title = r.FormValue("title")
	book.Description = r.FormValue("desc")
	db.Save(&book)
	// fmt.Println(book)
	var response = m.JsonResponse{}
	if id == "" || book.Title == "" || book.Description == "" {
        response = m.JsonResponse{Type: "error", Message: "You are missing id or title or description parameter."}
	} else if book.Id != ""{
		response = m.JsonResponse{Type: "success", Message: "Updated!" , Data: book}
	} else {
		response = m.JsonResponse{Type: "error", Message: "Not found!"}
	}
	json.NewEncoder(w).Encode(response)
}

func DeleteById(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	var response = m.JsonResponse{}
	book := m.Book{}
	db.Where("id=? and id != '' ", id).Find(&book)
	deleted := book
	db.Delete(&book, id)
	// fmt.Println(book)
	if id == "" {
        response = m.JsonResponse{Type: "error", Message: "You are missing id parameter."}
    } else if deleted.Id!=""{
		response = m.JsonResponse{Type: "success", Message: "Deleted!"}
	} else {
		response = m.JsonResponse{Type: "error", Message: "Not found!"}
	}
	json.NewEncoder(w).Encode(response)
}

func SearchByTitle(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	var response = m.JsonResponse{}
	book := []m.Book{}
	db.Where("title = ?", title).Find(&book)
	if title == "" {
        response = m.JsonResponse{Type: "error", Message: "You are missing title parameter."}
    } else if len(book) !=0 {
		response = m.JsonResponse{Type: "success", Message: "Found!" , Datas: book}
	} else {
		response = m.JsonResponse{Type: "error", Message: "Not found!"}
	}
	json.NewEncoder(w).Encode(response)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	title := r.FormValue("title")
	description := r.FormValue("desc")
	cost:=r.FormValue("cost")

	c, err := strconv.Atoi(cost)
    if err != nil {
        fmt.Println(err)
    }

	// db := DB()
	var response = m.JsonResponse{}
	if id == "" || title == "" || description == "" || c == 0{
        response = m.JsonResponse{Type: "error", Message: "You are missing id or title or description or cost parameter."}
    } else {
		db.Create(&m.Book{Id: id, Title: title, Description: description, Cost: c})
		response = m.JsonResponse{Type: "success", Message: "Added!"}
	}
	json.NewEncoder(w).Encode(response)
}

func SortOrderByAsc(w http.ResponseWriter, r *http.Request){
	var response = m.JsonResponse{}
	book := []m.Book{}
	db.Order("cost asc").Find(&book)
	// fmt.Println(book)
	response = m.JsonResponse{Type: "success", Message: "Books:", Datas: book}
	json.NewEncoder(w).Encode(response)
}

func SortOrderByDesc(w http.ResponseWriter, r *http.Request){
	var response = m.JsonResponse{}
	book := []m.Book{}
	db.Order("cost desc").Find(&book)
	// fmt.Println(book)
	response = m.JsonResponse{Type: "success", Message: "Books:", Datas: book}
	json.NewEncoder(w).Encode(response)
}