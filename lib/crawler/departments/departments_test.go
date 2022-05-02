package departments

import (
	"io/ioutil"
	"testing"
)

func TestDepartments(t *testing.T) {
	deps := GetDepartments()
	if len(deps) < 1 {
		t.Error("error getting amazon departments")
	}else {
		data := JsonMarshal(deps)
		err := ioutil.WriteFile("/home/ubuntu/Documents/amazon/testData/deps.json", data, 0755)
		HandleError(err)
	}
}