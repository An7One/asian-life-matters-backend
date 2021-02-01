package controller

import (
	"encoding/json"
	"log"
	"net/http"

	user "github.com/zea7ot/web_api_aeyesafe/model"
)

func registerUserRoutes() {
	// http.HandleFunc("/users/login", handleUserLogin)
	http.HandleFunc("/user/signup", handleUserSignUp)
}

// func handleUserLogin(w http.ResponseWriter, r *http.Request) {
// 	pattern, _ := regexp.Compile(`/users/(\d+)`)
// 	matches := pattern.FindStringSubmatch(r.URL.Path)
// 	if len(matches) > 0 {
// 		userId, _ := strconv.Atoi(matches[1])
// 	} else {

// 	}
// }

func handleUserSignUp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		dec := json.NewDecoder(r.Body)
		user := user.UserSignUp{}
		err := dec.Decode(&user)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		enc := json.NewEncoder(w)
		err = enc.Encode(user)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
