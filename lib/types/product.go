package types

import (

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID										primitive.ObjectID					`json:"_id,omitempty" bson:"_id,omitempty"`
	Asin									string								`json:"asin,omitempty" bson:"asin,omitempty"`
	Title									string								`json:"title,omitempty" bson:"title,omitempty"`
	PriceRange								[]Price								`json:"price,omitempty" bson:"price,omitempty"`
	Features								[]Feature							`json:"features,omitempty" bson:"features,omitempty"`
	FeatureBullets							[]string							`json:"feature-bullets,omitempty" bson:"feature-bullets"`
	Sizes									[]string							`json:"sizes,omitempty" bson:"sizes,omitempty"`
	Images									[]Image								`json:"images,omitempty" bson:"images,omitempty"`
	Colors									[]Color								`json:"colors,omitempty" bson:"colors,omitempty"`
	Videos									[]Video								`json:"videos,omitempty" bson:"videos,omitempty"`
	Details									Details								`json:"details,omitempty" bson:"details,omitempty"`
	Ratings									Ratings								`json:"ratings,omitempty" bson:"ratings,omitempty"`
	Description								string								`json:"description,omitempty" bson:"description,omitempty"`
	About									[]string							`json:"about,omitempty" bson:"about,omitempty"`
}

type Image struct {
	Tiny									string								`json:"tiny,omitempty" bson:"tiny,omitempty"`
	Normal									string								`json:"normal,omitempty" bson:"normal,omitempty"`
}

type Color struct {
	Asin									string								`json:"asin,omitempty" bson:"asin,omitempty"`
	Name									string								`json:"name,omitempty" bson:"name,omitempty"`
	Images									[]Image								`json:"images,omitempty" bson:"images,omitempty"`
}

type Video struct {
	Hls										string								`json:"hls,omitempty" bson:"hls,omitempty"`
	Mp4										string								`json:"mp4,omitempty" bson:"mp4,omitempty"`
}

type Details struct {
	PackageDimensions						string								`json:"package-dimensions,omitempty" bson:"package-dimensions,omitempty"`
	DatePublished							string								`json:"date-published,omitempty" bson:"date-published,omitempty"`
	Manufacture								string								`json:"manufacture,omitempty" bson:"manufacture,omitempty"`
	Weight									string								`json:"weight,omitempty" bson:"weight,omitempty"`
	Brand									string								`json:"brand,omitempty" bson:"brand,omitempty"`
	ManufacturePartNumber					string								`json:"manufacture-part-number,omitempty" bson:"manufacture-part-number,omitempty"`
	ModelNumber								string								`json:"model-number,omitempty" bson:"model-number"`
	IsDiscontinued							bool								`json:"discontinued,omitempty" bson:"discontinued,omitempty"`
	Origin									string								`json:"origin,omitempty" bson:"origin,omitempty"`
	Color									string								`json:"color,omitempty" bson:"color,omitempty"`
	Asin									string								`json:"asin,omitempty" bson:"asin,omitempty"`
}

type Star struct {
	Number									int									`json:"number,omitempty" bson:"number,omitempty"`
	Percentage								int									`json:"percentage,omitempty" bson:"percentage,omitempty"`
}

type Feature struct {
	Key										string								`json:"key,omitempty" bson:"key,omitempty"`
	Value									string								`json:"value,omitempty" bson:"value,omitempty"`
}

type Ratings struct {
	TotalStars								float64								`json:"total-stars,omitempty" bson:"total-stars,omitempty"`
	RatedStars								float64								`json:"rated-stars,omitempty" bson:"rated-stars,omitempty"`
	Stars									[]Star								`json:"stars,omitempty" bson:"stars,omitempty"`
	TotalReviews							int									`json:"total-reviews,omitempty" bson:"total-reviews"`
	ReviewsUrl								string								`json:"review-url,omitempty" bson:"review-url,omitempty"`
}

type Price struct {
	Currency								string								`json:"currency,omitempty" bson:"currency,omitempty"`
	Price									float64								`json:"price,omitempty" bson:"price,omitempty"`
}