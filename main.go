package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", Index).Methods("GET")
	r.HandleFunc("/todo", getAllTasks).Methods("GET")
	r.HandleFunc("/todo/{id}", getSpecificTask).Methods("GET")

	http.Handle("/", r)
	fmt.Println("= server will start =")
	log.Fatal(http.ListenAndServe(":5002", r))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HELLO WORLD")
}

func getAllTasks(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	// q := Todo{gorm.Model.Deleted_at: nil}
	todo := []models.Todo{}
	rr := db.Find(&todo)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rr.Value)
}

func getSpecificTask(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	// id, _ := strconv.Atoi(params["id"])

	// for _, v := range tasks {
	// 	if v.Id == id {
	// 		w.Header().Set("Content-Type", "application/json")
	// 		json.NewEncoder(w).Encode(v)
	// 		fmt.Println(v)
	// 		return
	// 	}
	// }
	// w.WriteHeader(400)
	// json.NewEncoder(w).Encode([]string{})
}
