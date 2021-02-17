package model

import (
	"math/rand"
	"strconv"
	"time"
)

// ProfileOTP the model for one-time password
type ProfileOTP struct {
	PhoneNumber string `json:"phoneNumber"`

	// one time password
	OTP          string    `json:"otp"`
	OTPCreatedAt time.Time `json:"otpCreatedAt"`
}

// GenerateOTP generates a one-time password， consisting of 7 pure digits
func GenerateOTP() string {
	return strconv.Itoa(rangeIn(1000000, 9999999))
}

func rangeIn(low, high int) int {
	return low + rand.Intn(high-low)
}
