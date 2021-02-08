package twilio

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/spf13/viper"
)

// Message carries messages to be sent via Twilio
type Message struct {
	PhoneNumberTo  string
	MessageContent string
}

// SendMessage sends SMS messages via Twilio
func (mes *Message) SendMessage() {
	accountSid := viper.GetString("twilio.account_sid")
	twilioAuthToken := viper.GetString("twilio.auth_token")
	phoneNumberDefaultFrom := viper.GetString("phone_number_default_from")
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	msgData := url.Values{}
	msgData.Set("To", mes.PhoneNumberTo)
	msgData.Set("From", phoneNumberDefaultFrom)
	msgData.Set("Body", mes.MessageContent)
	msgDataReader := *strings.NewReader(msgData.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, twilioAuthToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err != nil {

		}
	} else {
		// to deal with unsuccessful POST request to Twilio API
	}
}
