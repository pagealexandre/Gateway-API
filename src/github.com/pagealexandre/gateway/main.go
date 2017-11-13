package main

import (
	"os"
	"fmt"
	"net/http"
    "github.com/gorilla/mux"
	"github.com/nats-io/nats"
    "github.com/pagealexandre/gateway/handler"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Wrong number of arguments. Need NATS server address.")
        return
    }
    var err error
    var nc *nats.Conn

    nc, err = nats.Connect(os.Args[1])
    if err != nil {
        fmt.Println(err)
    }

    m := mux.NewRouter()
    user := m.PathPrefix("/user").Subrouter()
    user.HandleFunc("/{uuid}", handler.GetUser(nc)).Methods("GET")
    user.HandleFunc("/", handler.RegistrationUser(nc)).Methods("POST")

    list := m.PathPrefix("/list").Subrouter()
    list.HandleFunc("/", handler.CreationList(nc)).Methods("POST")
    list.HandleFunc("/{id}", handler.GetList(nc)).Methods("GET")
    list.HandleFunc("/{id}", handler.DeleteList(nc)).Methods("DELETE")
    list.HandleFunc("/", handler.UpdateList(nc)).Methods("PUT")

    task := m.PathPrefix("/task").Subrouter()
    task.HandleFunc("/", handler.CreationTask(nc)).Methods("POST")
    task.HandleFunc("/{id}", handler.GetTask(nc)).Methods("GET")
    task.HandleFunc("/{id}", handler.DeleteTask(nc)).Methods("DELETE")
    task.HandleFunc("/", handler.UpdateTask(nc)).Methods("PUT")

    http.ListenAndServe(":3000", m)
}

