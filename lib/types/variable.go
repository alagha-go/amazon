package types

import (
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gorilla/mux"
)

var (
	Client *mongo.Client
	Router = mux.NewRouter()
	CouldNotRetrieveData = Error{Error: "could not retrieve data from the database."}
	CouldNotInsertData = Error{Error: "could not insert data to the database."}
	UnAuthorizedAcces = Error{Error: "un Authorized Access."}
	ExpiredToken = Error{Error: "expired Authorization token."}
	InvalidApiKey = Error{Error: "invalid api key."}
	InvalidID = Error{Error: "invalid id."}
	PaymentExpired = Error{Error: "Your payment has expired. Update your payment to access this service."}
	UserDoesNotExist = Error{Error: "User does not exist."}
	DataDoesNotExist = Error{Error: "could not find the requested data."}
)