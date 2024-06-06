package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

var users []User

func init() {
	users = []User{{
		ID:        1,
		FirstName: "John",
		LastName:  "Smith",
		Email:     "john.smith@gmail.com",
	}, {ID: 2,
		FirstName: "Jane",
		LastName:  "Smith",
		Email:     "jane.smith@gmail.com",
	}, {ID: 3,
		FirstName: "Jan88e",
		LastName:  "Smith",
		Email:     "jan88e.smith@gmail.com",
	},
	}
}
func main() {

	http.HandleFunc("/users", UserService)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
func UserService(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetAllUser(w)

		//w.WriteHeader(http.StatusOK)
		//fmt.Fprintf(w, `{"status": %d, "message": "%s"}`, http.StatusOK, "SUCCESS in get ")
	case http.MethodPost:
		decode := json.NewDecoder(r.Body)
		var user User
		if err := decode.Decode(&user); err != nil {
			MsgResponse(w, http.StatusInternalServerError, err.Error())
			return

		}
		users = append(users, user)
		DataResponse(w, http.StatusCreated, user)
		//w.WriteHeader(http.StatusCreated)
		//fmt.Fprintf(w, `{"status": %d, "message": "%s"}`, http.StatusCreated, "SUCCESS in post ")
	case http.MethodPut:
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, `{"status": %d, "message": "%s"}`, http.StatusAccepted, "SUCCESS in put ")
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"status": %d, "message": "%s"}`, http.StatusBadRequest, "not found ")
	}
}
func GetAllUser(w http.ResponseWriter) {
	DataResponse(w, http.StatusOK, users)

}
func MsgResponse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "mesage": %s}`, status, message)
}
func DataResponse(w http.ResponseWriter, status int, users interface{}) {
	value, err := json.Marshal(users)
	if err != nil {
		MsgResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "data": %s}`, status, value)
}
