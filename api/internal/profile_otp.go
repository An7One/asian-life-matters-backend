package app

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	validation "github.com/go-ozzo/ozzo-validation"
	model "github.com/zea7ot/web_api_aeyesafe/model/user"
)

// the list of errors returned from the profile resource
var (
	ErrProfileOTPValidation = errors.New("Profile One-Time Password validation error")
)

// ProfileOTPDBClient defines database operations for a ProfileOTP
type ProfileOTPDBClient interface {
	AddOneProfileOTP(otp *model.ProfileOTP) (*model.ProfileOTP, error)
	GetOneProfileOTPByPhoneNumber(phoneNumber string) (*model.ProfileOTP, error)
	UpdateOneProfileOTP(otp *model.ProfileOTP) (*model.ProfileOTP, error)
}

// ProfileOTPResource implements ProfileOTP management resource
type ProfileOTPResource struct {
	Client ProfileOTPDBClient
}

// NewProfileOTPResource creates and returns a ProfileOTP resource
func NewProfileOTPResource(client ProfileOTPDBClient) *ProfileOTPResource {
	return &ProfileOTPResource{
		Client: client,
	}
}

func (rs *ProfileOTPResource) router() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", rs.get)
	r.Post("/verify", rs.verify)
	r.Put("/", rs.update)

	return r
}

func (rs *ProfileOTPResource) get(w http.ResponseWriter, r *http.Request) {
	var otp *model.ProfileOTP
	err := json.NewDecoder(r.Body).Decode(&otp)
	if err != nil {
		render.Render(w, r, ErrValidation(ErrProfileOTPValidation, err.(validation.Errors)))
		return
	}

	res, err := rs.Client.GetOneProfileOTPByPhoneNumber(otp.PhoneNumber)
	if err != nil {
		switch err.(type) {
		case validation.Errors:
			render.Render(w, r, ErrValidation(ErrProfileOTPValidation, err.(validation.Errors)))
			return
		}
		render.Render(w, r, ErrRender(err))
		return
	}

	render.Respond(w, r, newProfileOTPResponse(res))
}

func (rs *ProfileOTPResource) update(w http.ResponseWriter, r *http.Request) {

}

func (rs *ProfileOTPResource) verify(w http.ResponseWriter, r *http.Request) {
	var otp *model.ProfileOTP
	err := json.NewDecoder(r.Body).Decode(&otp)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	extOtp, err := rs.Client.GetOneProfileOTPByPhoneNumber(otp.PhoneNumber)
	if err != nil {
		switch err.(type) {
		case validation.Errors:
			render.Render(w, r, ErrValidation(ErrProfileOTPValidation, err.(validation.Errors)))
			return
		}
		render.Render(w, r, ErrRender(err))
		return
	}

	if otp.OTP != extOtp.OTP {
		// 403 not authorized

	} else {
		render.Respond(w, r, newProfileOTPResponse(otp))
	}
}

type profileOTPResponse struct {
	*model.ProfileOTP
}

func newProfileOTPResponse(otp *model.ProfileOTP) *profileOTPResponse {
	return &profileOTPResponse{
		ProfileOTP: otp,
	}
}
