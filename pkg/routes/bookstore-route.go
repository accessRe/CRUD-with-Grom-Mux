package routes

import (
	"github.com/accessre/go/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

// Creating Routes and linking them to controllers for a specific function like createbook etc...
var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{ID}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{Id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{book{Id}", controllers.DeleteBook).Methods("DELETE")

}
