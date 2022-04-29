package handler

import (
	"io/ioutil"
	"net/http"
	"log"
)


func GetAllDepartments(res http.ResponseWriter, req *http.Request) {
	data, err := ioutil.ReadFile("/home/ubuntu/Documents/amazon/testData/deps.json")
	HandleError(err)
	res.WriteHeader(200)
	res.Header().Set("content-type", "application/json")
	res.Write(data)
}

func HandleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}