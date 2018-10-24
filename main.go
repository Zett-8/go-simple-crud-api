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
	r.HandleFunc("/todo", getTasks).Methods("GET")
	r.HandleFunc("/todo/create/", createTask).Methods("POST")
	r.HandleFunc("/todo/{id}", getTaskById).Methods("GET")
	r.HandleFunc("/todo/{id}/update/", updateTask).Methods("PUT")
	r.HandleFunc("/todo/{id}/delete/", deleteTask).Methods("DELETE")

	http.Handle("/", r)
	fmt.Println("= server will start =")
	log.Fatal(http.ListenAndServe(":5002", r))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HELLO WORLD")
}

func getTasks(w http.ResponseWriter, r *http.Request) {
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
	todo_id, _ := strconv.Atoi(params["id"])

	todo := models.Todo{}
	res := db.First(&todo, todo_id)

	if res.Error != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode([]string{})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res.Value)
	return
}

func createTask(w http.ResponseWriter, r *http.Request) {

	db := utils.ConnectDB()
	defer db.Close()

	var newTodo models.Todo

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newTodo)
	if err != nil {
		panic(err)
	}

	if newTodo.Name == "" {
		w.WriteHeader(400)
		return
	}

	db.Create(&newTodo)

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(newTodo)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	db := utils.ConnectDB()
	defer db.Close()

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	oldTodo := models.Todo{}
	oldTodo.ID = uint(id)

	newTodo := models.Todo{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newTodo)
	if err != nil {
		panic(err)
	}

	if newTodo.Name == "" {
		w.WriteHeader(400)
		return
	}

	db.Model(&oldTodo).Update(&newTodo)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {

}
