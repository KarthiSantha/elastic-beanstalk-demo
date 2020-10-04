package main
import (
    "github.com/gorilla/mux"
	"log"
	"net/http"

	"beanstalk_go_app/controllers"
	
) 

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/hello/{name}/{age}/{email}",controller.HelloWorld)

	log.Print("Listening on port 8080")

	log.Fatal(http.ListenAndServe(":8080", router))

}



