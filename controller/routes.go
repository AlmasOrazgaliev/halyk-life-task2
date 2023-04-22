package controller

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Start(controller *Controller) error {
	router := mux.NewRouter()
	router.HandleFunc("/authors", controller.HandleAuthors)
	router.HandleFunc("/books", controller.HandleBooks)
	router.HandleFunc("/members", controller.HandleMembers)

	router.HandleFunc("/authors/{id:[0-9]+}", controller.HandleAuthors)
	router.HandleFunc("/books/{id:[0-9]+}", controller.HandleBooks)
	router.HandleFunc("/members/{id:[0-9]+}", controller.HandleMembers)

	router.HandleFunc("/members/{memberId:[0-9]+}/subscribe", controller.HandleSubscribe)

	router.HandleFunc("/authors/{id:[0-9]+}/books", controller.HandleAuthorBooks)
	router.HandleFunc("/members/{id:[0-9]+}/books", controller.HandleMemberBooks)
	return http.ListenAndServe(":8080", router)
}
