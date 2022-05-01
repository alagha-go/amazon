package handler

import (
	"net/http"
	"log"
)


func Main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/departments", GetAllDepartments)
}


func Hello(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("<h1>Hello World!!!</h1>"))
}

func HandleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}