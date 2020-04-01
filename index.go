package main

import (
	"fmt"
	"net/http"

	"crud_tutor/config"
	"crud_tutor/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
)

func main() {
	fmt.Printf(config.Testing() + "\n")

	router := mux.NewRouter()
	router.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ini Produk"))
	})
	router.HandleFunc("/employees", controllers.AllEmployees).Methods("GET")
	router.HandleFunc("/employees", controllers.AddEmployees).Methods("POST")
	router.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)
	http.Handle("/", router)

	address := "localhost:9000"
	fmt.Printf("server started at %s\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	message := "Welcome"
	w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	message := "Hello world!"
	w.Write([]byte(message))
}
