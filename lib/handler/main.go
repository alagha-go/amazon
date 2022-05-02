package handler

import (
	"amazon/lib/types"
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func Main() {
	var Router *mux.Router = types.Router
	Router.HandleFunc("/", Hello)
	Router.HandleFunc("/departments", GetAllDepartments)
}


func Hello(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("<h1>Hello World!!!</h1>"))
}

func HandleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func JsonMarshal(data interface{}) []byte {
	var buff bytes.Buffer
	encoder := json.NewEncoder(&buff)
	encoder.SetEscapeHTML(false)
	encoder.Encode(data)
	return buff.Bytes()
}