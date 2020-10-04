package model

import (
	"net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}


type User struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:"email"`
}