package dataloaders

import (
	/*
		Nokta yerine m yazsaydık "m.Users" kullanımı olacaktı.
		Şimdi ise direkt "Users" olarak kullanabiliriz.
	*/
	"encoding/json"

	. "rest-api-kullanici-sistemi-ornegi/models"
	utils "rest-api-kullanici-sistemi-ornegi/utils"
)

// LoadUsers ...
// JSON formatındaki veriyi Users nesnesine dönüştüren fonksiyon.
func LoadUsers() []Users {
	bytes, _ := utils.ReadFile("../json/users.json")
	var users []Users
	json.Unmarshal([]byte(bytes), &users)
	return users
}

// LoadInterests ...
func LoadInterests() []Interest {
	bytes, _ := utils.ReadFile("../json/interests.json")
	var interests []Interest
	json.Unmarshal([]byte(bytes), &interests)
	return interests
}

// LoadInterestMappings ...
func LoadInterestMappings() []InterestMapping {
	bytes, _ := utils.ReadFile("../json/userInterestMappings.json")
	var data []InterestMapping
	json.Unmarshal([]byte(bytes), &data)
	return data
}
