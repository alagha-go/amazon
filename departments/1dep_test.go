package departments

import (
	"bytes"
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
		data := JsonMarshal(deps)
		err := ioutil.WriteFile("../DB/deps.json", data, 0755)
		HandleError(err)
	}
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