package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm"
	"net/http"
	"time"
	_ "time"
)

var db = connect()

type Service struct {
	Name            string `json:"name"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}

func index (w http.ResponseWriter, r *http.Request) {
	var service []Service
	db.Find(&service)

	json.NewEncoder(w).Encode(service)
}

func show (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var service []Service
	db.Where("id = ?", id).Find(&service)
	json.NewEncoder(w).Encode(service)
}

func create (w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	name := r.FormValue("name")
	db.Create(&Service{
		Name:            name,
		CreatedAt:       time.Now().UTC(),
		UpdatedAt:       time.Now().UTC(),
	})
	fmt.Fprintln(w, "New service created successfully!")
}

func update (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var service Service
	db.Where("id = ?", id).Find(&service)

	name := r.FormValue("name")

	db.Model(&service).Where("id = ?", id).Update(&Service{
		Name:      name,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	})

	fmt.Fprintln(w, "Service updated successfully!")
}

func delete (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var service Service

	db.Where("id = ?", id).Delete(&service)
	fmt.Fprintln(w, "Service delete successfully!")

}