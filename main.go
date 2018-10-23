package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Zett-8/go-simple-crud-api/models"
	"github.com/Zett-8/go-simple-crud-api/utils"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", Index).Methods("GET")
	r.HandleFunc("/todo", getAllTasks).Methods("GET")
	r.HandleFunc("/todo/{id}", getTaskById).Methods("GET")

	http.Handle("/", r)
	fmt.Println("= server will start =")
	log.Fatal(http.ListenAndServe(":5002", r))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HELLO WORLD")
}

func getAllTasks(w http.ResponseWriter, r *http.Request) {
	db := utils.ConnectDB()
	defer db.Close()

	todo := []models.Todo{}
	res := db.Find(&todo)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res.Value)
}

func getTaskById(w http.ResponseWriter, r *http.Request) {
	db := utils.ConnectDB()
	defer db.Close()

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	todo := models.Todo{}
	res := db.First(&todo, id)

	if res.Error != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode([]string{})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
	return
}
