package crawler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"
)

func TestDepartments(t *testing.T) {
	deps := GetDepartments()
	if len(deps) < 1 {
		t.Error("error getting amazon departments")
	}else {
		data, err := json.Marshal(deps)
		HandleError(err)
		err = ioutil.WriteFile("/home/ubuntu/Documents/amazon/testData/deps.json", data, 0755)
		HandleError(err)
	}
}

func HandleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}