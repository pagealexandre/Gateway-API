package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/nats-io/nats"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
	"github.com/pagealexandre/gateway/transport"
)

func GetUser(nc *nats.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		user := transport.User{Uid: vars["uuid"]}

		data, err := proto.Marshal(&user)
		if err != nil {
			fmt.Println("Problem parsing uid")
			return
		}

		msg, err := nc.Request("service.user.get", data, 10*time.Second) // Increase time out if pannic error
		err = proto.Unmarshal(msg.Data, &user)
		if err != nil {
			fmt.Println(err)
		} else {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(user); err != nil {
				panic(err)
			}
		}		
	}
}

func RegistrationUser(nc *nats.Conn) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		userName := r.FormValue("userName")
		prettyName := r.FormValue("prettyName")
		password := r.FormValue("password")

		if userName == "" || prettyName == "" || password == "" {
			if err := json.NewEncoder(w).Encode("error"); err != nil {
				panic(err)
			}
		}

		user := transport.User{UserName: userName, PrettyName: prettyName, Password: password}

		data, err := proto.Marshal(&user)

		msg, err := nc.Request("service.user.register", data, 1*time.Second) // Increase time out if pannic error
		err = proto.Unmarshal(msg.Data, &user)

		if err != nil {
			fmt.Println(err)
		} else {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(user); err != nil {
				panic(err)
			}
		}		
	}

}