package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

// Profile the model for user
type Profile struct {
	ID int `json:"-"`
	// Username    string `json:"username"`
	PhoneNumber string `json:"phonenumber"`
	// Password    string `json:"password,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`

	Theme string `json:"theme,omitempty"`
}

// BeforeInsert hooks executed before database insersion operation
func (p *Profile) BeforeInsert() error {
	p.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate hooks executed before database update operation
func (p *Profile) BeforeUpdate() error {
	p.UpdatedAt = time.Now()
	return p.Validate()
}

// Validate validates Profile struct and returns validation errors
func (p *Profile) Validate() error {
	return validation.ValidateStruct(p,
		validation.Field(&p.Theme, validation.Required, validation.In("default", "dark")),
	)
}
