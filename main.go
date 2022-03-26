package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Student struct {
	Name    string `json:"Name"`
	Address string `json:"Address"`
	Marks   int    `json:"Marks"`
}

type Students []Student

func allStudents(w http.ResponseWriter, r *http.Request) {
	students := Students{
		Student{Name: "pravin prabhu", Address: "pune", Marks: 70},
		{Name: "Harshit Srivastava", Address: "Bhopal", Marks: 80},
		{Name: "Surabhi Singh", Address: "Ranchi", Marks: 90},
		{Name: "Ran Malviya", Address: "Harda", Marks: 60},
		{Name: "Roohi Yadav", Address: "Madhubani", Marks: 65},
	}
	fmt.Println("Printing All Students")
	json.NewEncoder(w).Encode(students)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi My Name is Pravin Prabhu")
}

func handleRequests() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/students", allStudents)
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func main() {
	handleRequests()
}
