package app

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	model "github.com/zea7ot/web_api_aeyesafe/model/user"
)

var (
	ErrProfileValidation = errors.New("profile validation error")
)

// ProfileStore defines database operations for a profile
type ProfileStore interface {
	Get(phonenumber string) (*model.Profile, error)
	Update(p *model.Profile) error
}

// ProfileResource implements profile management handler
type ProfileResource struct {
	Store ProfileStore
}

// NewProfileResource creates and returns a profile resource
func NewProfileResource(store ProfileStore) *ProfileResource {
	return &ProfileResource{
		Store: store,
	}
}

func (rs *ProfileResource) router() *chi.Mux {
	r := chi.NewRouter()
	r.Use(rs.profileCtx)
}

func (rs *ProfileResource) profileCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
