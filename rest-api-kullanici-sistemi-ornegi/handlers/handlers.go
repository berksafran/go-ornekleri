package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	dl "rest-api-kullanici-sistemi-ornegi/dataloaders"
	. "rest-api-kullanici-sistemi-ornegi/models"
)

// Run ...
func Run() {
	http.HandleFunc("/", Handler)
	fmt.Println("Listening on 8080 port...")
	http.ListenAndServe(":8080", nil)
}

// Handler ...
func Handler(w http.ResponseWriter, r *http.Request) {
	page := Page{
		ID:          7,
		Name:        "Kullanıcılar",
		Description: "Kullanıcı Listesi",
		URI:         "/users",
	}

	users := dl.LoadUsers()
	interests := dl.LoadInterests()
	interestMappings := dl.LoadInterestMappings()

	var newUsers []Users

	for _, user := range users {
		for _, interestMapping := range interestMappings {
			if user.ID == interestMapping.UserID {
				for _, interest := range interests {
					if interestMapping.InterestID == interest.ID {
						user.Interests = append(user.Interests, interest)
					}
				}
			}
		}
		newUsers = append(newUsers, user)
	}
	viewModel := UserViewModel{Page: page, Users: newUsers}
	data, _ := json.Marshal(viewModel)
	w.Write([]byte(data))
}
