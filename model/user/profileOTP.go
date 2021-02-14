package model

import (
	"time"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/sirupsen/logrus"
)

// ProfileOTP the model for one-time password
type ProfileOTP struct {
	PhoneNumber string `json:"phoneNumber"`

	// one time password
	OTP          string    `json:"otp"`
	OTPCreatedAt time.Time `json:"otpCreatedAt"`
}

// GenerateOTP generates a one-time password
func GenerateOTP() (string, error) {
	u4, err := uuid.NewV4()
	UUIDToken := u4.String()
	if err != nil {
		logrus.Errorln("error", err)
		return "", err
	}
	return UUIDToken, err
}
