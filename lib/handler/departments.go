package handler

import (
	"io/ioutil"
	"net/http"
)


func GetAllDepartments(res http.ResponseWriter, req *http.Request) {
	data, err := ioutil.ReadFile("/home/ubuntu/Documents/amazon/testData/deps.json")
	HandleError(err)
	res.WriteHeader(200)
	res.Header().Set("content-type", "application/json")
	res.Write(data)
}