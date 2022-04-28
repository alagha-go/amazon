package handler

import "net/http"


func Main() {
	http.HandleFunc("/", Hello)
}


func Hello(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("<h1>Hello World!!!</h1>"))
}