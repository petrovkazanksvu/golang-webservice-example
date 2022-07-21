package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/insanrizky/golang-webservice-example/controller"
	// "database/sql"
	// "github.com/go-sql-driver/mysql"
)

func home_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Go is super easy!")
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Contacts page!")
}

func main() {
	fmt.Println("Connecting to db")
	http.HandleFunc("/my/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.HandleFunc("/", user.SayHelloName)
	http.HandleFunc("/insert-user", user.InsertUser)
	http.HandleFunc("/login", user.Login)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
