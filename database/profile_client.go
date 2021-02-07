package database

import (
	model "github.com/zea7ot/web_api_aeyesafe/model/user"
)

type ProfileClient struct {
	client DBClient
}

// NewProfileClient returns a ProfileClient implementation
func NewProfileClient(client *DBClient) *ProfileClient {
	return &ProfileClient{
		client: client,
	}
}

func (c *ProfileClient) Get(phoneNumber string) (*model.Profile, error){
	p := model.Profile{PhoneNumber: phoneNumber}
	_, err := c.
}