package departments

import (
	"amazon/lib/types"
	"context"
)


func InsertOneDepartment(department types.Department) error {
	ctx := context.Background()
	collection := types.Client.Database("Amazon").Collection("Departments")

	_, err := collection.InsertOne(ctx, department)
	return err
}