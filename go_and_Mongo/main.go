package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/golang-web-dev/go_and_Mongo/jason/models"
	"net/http"
)

func main() {
	r := httprouter.New()
	r.GET("/", index)

	http.ListenAndServe(":8080", r)
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello, welcome!")
}

func getUser(w http.ResponseWriter, req *http.Request, p httprouter.Params){
	u:=

}
