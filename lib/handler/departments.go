package handler

import (
	"amazon/lib/departments"
	"net/http"
)


func GetAllDepartments(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	if req.Method != "GET" {
		res.WriteHeader(400)
		return
	}
	data, status := departments.GetDepartments()
	res.WriteHeader(status)
	res.Write(data)
}