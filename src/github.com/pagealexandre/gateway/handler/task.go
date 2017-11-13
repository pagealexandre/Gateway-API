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

func GetTask(nc *nats.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		task := transport.Task{Id: vars["id"]}

		data, err := proto.Marshal(&task)
		if err != nil {
			fmt.Println("Problem parsing uid")
			return
		}

		msg, err := nc.Request("service.task.retrieve", data, 10*time.Second) // Increase time out if pannic error
		err = proto.Unmarshal(msg.Data, &task)
		if err != nil {
			fmt.Println(err)
		} else {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(task); err != nil {
				panic(err)
			}
		}
	}
}

func CreationTask(nc *nats.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.FormValue("title")
		content := r.FormValue("content")
		idList := r.FormValue("idList")

		if title == "" || content == "" || idList == "" {
			if err := json.NewEncoder(w).Encode("error"); err != nil {
				panic(err)
			}
		}

		task := transport.Task{Title: title, Content: content, IdList: idList}

		data, err := proto.Marshal(&task)

		msg, err := nc.Request("service.task.create", data, 2*time.Second) // Increase time out if pannic error
		err = proto.Unmarshal(msg.Data, &task)

		if err != nil {
			fmt.Println(err)
		} else {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(task); err != nil {
				panic(err)
			}
		}
	}
}

func DeleteTask(nc *nats.Conn) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		if id == "" {
			if err := json.NewEncoder(w).Encode("error"); err != nil {
				panic(err)
			}
		}

		task := transport.Task{Id: id}

		data, err := proto.Marshal(&task)

		msg, err := nc.Request("service.task.delete", data, 10*time.Second) // Increase time out if pannic error
		err = proto.Unmarshal(msg.Data, &task)

		if err != nil {
			fmt.Println(err)
		} else {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(task); err != nil {
				panic(err)
			}
		}
	}
}

func UpdateTask(nc *nats.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.FormValue("title")
		content := r.FormValue("content")
		idList := r.FormValue("idList")
		id := r.FormValue("id")

		task := transport.Task{Id: id, Title: title, Content: content, IdList: idList}

		data, err := proto.Marshal(&task)
		if err != nil {
			fmt.Println("Problem parsing uid")
			return
		}

		msg, err := nc.Request("service.task.update", data, 10*time.Second) // Increase time out if pannic error
		err = proto.Unmarshal(msg.Data, &task)
		if err != nil {
			fmt.Println(err)
		} else {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(task); err != nil {
				panic(err)
			}
		}
	}
}