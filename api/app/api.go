package app

import (
	"net/http"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"github.com/zea7ot/web_api_aeyesafe/database"
	"github.com/zea7ot/web_api_aeyesafe/logging"
)

type ctxKey int

const (
	ctxAccount ctxKey = iota
	ctxProfile
)

// API provides application resources and handlers
type API struct {
	Profile *ProfileResource
}

// NewAPI configures and returns the application API
func NewAPI(db *dynamodb.DynamoDB) (*API, error) {
	profileDBClient := database.NewProfileClient(db)
	profile := NewProfileResource(profileDBClient)

	api := &API{
		Profile: profile,
	}

	return api, nil
}

// Router provides the application routes
func (a *API) Router() *chi.Mux {
	r := chi.NewRouter()

	r.Mount("/profile", a.Profile.router())

	return r
}

func log(r *http.Request) logrus.FieldLogger {
	return logging.GetLogEntry(r)
}
