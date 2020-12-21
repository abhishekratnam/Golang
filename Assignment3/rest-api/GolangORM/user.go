package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	// "github.com/mattn/go-sqlite3"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Name  string
	Email string
}

func InitialMigration() {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to the database")
	}
	defer db.Close()
	db.AutoMigrate(&User{})
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Cound not connect to the database")
	}
	defer db.Close()
	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}
func NewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All NewUSers ENdpoint Hit")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete USer ENdpoint Hit")
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update User Endpoint HIt")
}
