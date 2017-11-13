package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/nats-io/nats"
	"github.com/golang/protobuf/proto"
	"github.com/pagealexandre/gateway/transport"
	"github.com/gorilla/mux"
)

func CreationList(nc *nats.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.FormValue("title")
		idBoard := r.FormValue("idBoard")

		if title == "" || idBoard == "" {
			if err := json.NewEncoder(w).Encode("error"); err != nil {
				panic(err)
			}
		}

		list := transport.List{Title: title, IdBoard: idBoard}

		data, err := proto.Marshal(&list)

		msg, err := nc.Request("service.list.create", data, 2*time.Second) // Increase time out if pannic error
		err = proto.Unmarshal(msg.Data, &list)

		if err != nil {
			fmt.Println(err)
		} else {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(list); err != nil {
				panic(err)
			}
		}
	}
}

func GetList(nc *nats.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		list := transport.List{Id: vars["id"]}

		data, err := proto.Marshal(&list)
		if err != nil {
			fmt.Println("Problem parsing uid")
			return
		}

		msg, err := nc.Request("service.list.retrieve", data, 10*time.Second) // Increase time out if pannic error
		err = proto.Unmarshal(msg.Data, &list)
		if err != nil {
			fmt.Println(err)
		} else {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(list); err != nil {
				panic(err)
			}
		}
	}
}

func DeleteList(nc *nats.Conn) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		if id == "" {
			if err := json.NewEncoder(w).Encode("error"); err != nil {
				panic(err)
			}
		}

		list := transport.List{Id: id}

		data, err := proto.Marshal(&list)

		msg, err := nc.Request("service.list.delete", data, 10*time.Second) // Increase time out if pannic error
		err = proto.Unmarshal(msg.Data, &list)

		if err != nil {
			fmt.Println(err)
		} else {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(list); err != nil {
				panic(err)
			}
		}
	}
}

func UpdateList(nc *nats.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.FormValue("title")
		id := r.FormValue("id")
		idBoard := r.FormValue("idBoard")

		list := transport.List{Id: id, Title: title, IdBoard: idBoard}

		data, err := proto.Marshal(&list)
		if err != nil {
			fmt.Println("Problem parsing uid")
			return
		}

		msg, err := nc.Request("service.list.update", data, 10*time.Second) // Increase time out if pannic error
		err = proto.Unmarshal(msg.Data, &list)
		if err != nil {
			fmt.Println(err)
		} else {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(list); err != nil {
				panic(err)
			}
		}
	}
}