package postcontroller

import (
	"go-web-native-crud/models/postmodel"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	posts := postmodel.GetAll()

	data := map[string]any{
		"posts": posts,
	}

	temp, err := template.ParseFiles("views/post/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {

}

func Edit(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

}
