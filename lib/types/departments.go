package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Department struct {
	ID								primitive.ObjectID								`json:"_id,omitempty" bson:"_id,omitempty"`
	Title							string											`json:"title,omitempty" bson:"title,omitempty"`
	Url								string											`json:"url,omitempty" bson:"url,omitempty"`
	Categories						[]Category										`json:"categories,omitempty" bson:"categories,omitempty"`
}

type Category struct {
	ID								primitive.ObjectID								`json:"_id,omitempty" bson:"_id,omitempty"`
	Title							string											`json:"title,omitempty" bson:"title,omitempty"`
	Url								string											`json:"url,omitempty" bson:"url,omitempty"`
	SubCategories					[]SubCategory									`json:"subcategories,omitempty" bson:"subcategories,omitempty"`
}

type SubCategory struct {
	ID								primitive.ObjectID								`json:"_id,omitempty" bson:"_id,omitempty"`
	Title							string											`json:"title,omitempty" bson:"title,omitempty"`
	Url								string											`json:"url,omitempty" bson:"url,omitempty"`
	Types							[]Type											`json:"types,omitempty" bson:"types,omitempty"`
}

type Type struct {
	ID								primitive.ObjectID								`json:"_id,omitempty" bson:"_id,omitempty"`
	Title							string											`json:"title,omitempty" bson:"title,omitempty"`
	Url								string											`json:"url,omitempty" bson:"url,omitempty"`
	SubTypes						[]SubType										`json:"subtypes,omitempty" bson:"subtypes,omitempty"`
}

type SubType struct {
	ID								primitive.ObjectID								`json:"_id,omitempty" bson:"_id,omitempty"`
	Title							string											`json:"title,omitempty" bson:"title,omitempty"`
	Url								string											`json:"url,omitempty" bson:"url,omitempty"`
}