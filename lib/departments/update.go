package departments

import (
	"amazon/lib/types"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)


func UpdateDepartment(Department types.Department) {
	ctx := context.Background()
	collection := types.Client.Database("Amazon").Collection("Departments")

	filter := bson.M{
        "_id": bson.M{
            "$eq": Department.ID, // check if bool field has value of 'false'
        },
    }

	update := bson.M{"$set": bson.M{
		"Categories": Department.Categories,
	}}

	collection.UpdateOne(ctx, filter, update)
}