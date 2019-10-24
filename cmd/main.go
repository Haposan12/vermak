package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"vermak2/controllers"
	"vermak2/middleware"
)

func main()  {
	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")

	router.Use(middleware.JwtAuthentication)

	port := os.Getenv("PORT")
	////fmt.Sprintf(port)
	if port == ""{
		port = "1234"
	}

	fmt.Println("Start API in port:" +port)
	//fmt.Println(port)
	err := http.ListenAndServe(":"+port, router)

	if err != nil{
		fmt.Print(err)
	}
}
