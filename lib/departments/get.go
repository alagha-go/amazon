package departments

import (
	"amazon/lib/types"
	"bytes"
	"context"
	"encoding/json"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)


func GetDepartments() ([]byte, int) {
	var departments []types.Department
	ctx := context.Background()
	collection := types.Client.Database("Amazon").Collection("Departments")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return JsonMarshal(types.CouldNotRetrieveData), 500
	}
	cursor.All(ctx, &departments)

	return JsonMarshal(departments), 200
}

func JsonMarshal(data interface{}) []byte {
	var buff bytes.Buffer
	encoder := json.NewEncoder(&buff)
	encoder.SetEscapeHTML(false)
	encoder.Encode(data)
	return buff.Bytes()
}

func HandleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}