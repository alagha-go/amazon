package crawler


type Department struct {
	Title							string											`json:"title,omitempty" bson:"title,omitempty"`
	Url								string											`json:"url,omitempty" bson:"url,omitempty"`
	Categories						[]Category										`json:"categories,omitempty" bson:"categories,omitempty"`
	Category						*Category										`json:"category,omitempty" bson:"category,omitempty"`
}

type Category struct {
	Title							string											`json:"title,omitempty" bson:"title,omitempty"`
	Url								string											`json:"url,omitempty" bson:"url,omitempty"`
	SubCategories					[]SubCategory									`json:"subcategories,omitempty" bson:"subcategories,omitempty"`
	SubCategory						*SubCategory									`json:"subcategory,omitempty" bson:"subcategory,omitempty"`
}

type SubCategory struct {
	Title							string											`json:"title,omitempty" bson:"title,omitempty"`
	Url								string											`json:"url,omitempty" bson:"url,omitempty"`
	Types							[]Type											`json:"types,omitempty" bson:"types,omitempty"`
	Type							*Type											`json:"type,omitempty" bson:"type,omitempty"`
}

type Type struct {
	Title							string											`json:"title,omitempty" bson:"title,omitempty"`
	Url								string											`json:"url,omitempty" bson:"url,omitempty"`
	SubTypes						[]SubType										`json:"subtypes,omitempty" bson:"subtypes,omitempty"`
	SubType							*SubType										`json:"subtype,omitempty" bson:"subtype,omitempty"`
}

type SubType struct {
	Title							string											`json:"title,omitempty" bson:"title,omitempty"`
	Url								string											`json:"url,omitempty" bson:"url,omitempty"`
}