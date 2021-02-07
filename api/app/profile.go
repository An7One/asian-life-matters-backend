package app

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	validation "github.com/go-ozzo/ozzo-validation"
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
	r.Get("/", rs.get)
	r.Put("/", rs.update)
	return r
}

func (rs *ProfileResource) profileCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

type profileRequest struct {
	*model.Profile
	ProtectedID int `json:"id"`
}

func (d *profileRequest) Bind(r *http.Request) error {
	return nil
}

type profileResponse struct {
	*model.Profile
}

func newProfileResponse(p *model.Profile) *profileResponse {
	return &profileResponse{
		Profile: p,
	}
}

func (rs *ProfileResource) get(w http.ResponseWriter, r *http.Request) {
	p := r.Context().Value(ctxProfile).(*model.Profile)
	render.Respond(w, r, newProfileResponse(p))
}

func (rs *ProfileResource) update(w http.ResponseWriter, r *http.Request) {
	p := r.Context().Value(ctxProfile).(*model.Profile)
	data := &profileRequest{Profile: p}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
	}

	if err := rs.Store.Update(p); err != nil {
		switch err.(type) {
		case validation.Errors:
			render.Render(w, r, ErrValidation(ErrProfileValidation, err.(validation.Errors)))
			return
		}
		render.Render(w, r, ErrRender(err))
		return
	}
	render.Respond(w, r, newProfileResponse(p))
}
