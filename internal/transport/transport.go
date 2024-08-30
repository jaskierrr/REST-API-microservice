package transport

import (
	"card-project/internal/services"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Run() {
	router := mux.NewRouter()

	router.HandleFunc("/", helloWorld)
	router.HandleFunc("/add/user", services.AddUser).Methods("POST")
	router.HandleFunc("/delete/user/{id}", services.DeleteUser).Methods("DELETE")
	router.HandleFunc("/update/user/{id}", services.UpdateUser).Methods("PUT")

	fmt.Println("server running on port: 8080")
	http.ListenAndServe(":8080", router)
}


func helloWorld(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write([]byte("Helo World!"))
	}
}
