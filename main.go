package main

import (
	"go-web-native-crud/config"
	"go-web-native-crud/controllers/categorycontroller"
	"go-web-native-crud/controllers/homecontroller"
	"log"
	"net/http"
)

func main() {

	config.ConnectDB()

	// Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	// Category
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	log.Println("Server up and Running")
	http.ListenAndServe(":8080", nil)
}