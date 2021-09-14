package reader

import (
	_ "github.com/gin-gonic/gin"
	_ "github.com/mercadolibre/fury_go-platform/pkg/fury"
)


func(){

}


	/*//Main -> Domain -> Service -> Repo.
	//endpoints
	router.HandleFunc("/users", handler.GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", handler.GetByID).Methods("GET")
	router.HandleFunc("/user", handler.CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", handler.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", handler.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", router))

}
