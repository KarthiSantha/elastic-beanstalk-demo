package controller

import (
    "log"
    "encoding/json"
	"strconv"
	"net/http"
	"github.com/gorilla/mux"

	"beanstalk_go_app/model"
)

func HelloWorld (w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	
	name := vars["name"]
	age,_ := strconv.Atoi(vars["age"])
	email := vars["email"]

	p := model.User{Name:name,Age:age,Email:email}
	w.Header().Set("Content-Type", "application/json")
	log.Print("Request received on port 8080")
	json.NewEncoder(w).Encode(p)
	
}