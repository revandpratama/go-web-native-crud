package main

import (
	"go-web-native-crud/config"
	"go-web-native-crud/controllers/categorycontroller"
	"go-web-native-crud/controllers/homecontroller"
	"go-web-native-crud/controllers/postcontroller"
	"go-web-native-crud/controllers/productcontroller"
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

	// Post
	http.HandleFunc("/posts", postcontroller.Index)
	http.HandleFunc("/posts/add", postcontroller.Add)
	http.HandleFunc("/posts/edit", postcontroller.Edit)
	http.HandleFunc("/posts/delete", postcontroller.Delete)

	// Product
	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)

	log.Println("Server up and Running")
	http.ListenAndServe(":8080", nil)
}
