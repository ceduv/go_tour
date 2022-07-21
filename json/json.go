package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name       string  `json:"name"` // `json:"key"`
	Age        int     `json:"age"`
	Email      string  `json:"email"`
	Phone      string  `json:"phone"`
	City       *string `json:"ville"`
	Country    string  `json:"country,omitempty"`
	PostalCode string  `json:"-"`
}

func main() {
	jsonFromApi := `
		[
			{
				"name": "cedric",
				"age": 34,
				"email": "ced@gmail.com",
				"phone": "0619770606",
				"ville": "Paris",
				"postalCode": "75009"
			},
			{
				"name": "alex",
				"age": 30,
				"email": "alex@gmail.com",
				"phone": "061977010",
				"postalCode": "75009"
			}
		]
	`
	fmt.Printf("\n---------👿👿👿----------\n")
	fmt.Printf("json before Unmarshall : %v\n", jsonFromApi)

	// passez de notre format json a un slice

	// on init un slice users de type User(structure)
	var users []User

	// gestion de l err (si il y a une err) et parsing dans users
	if err := json.Unmarshal([]byte(jsonFromApi), &users); err != nil {
		fmt.Println("Error unmarshalling json :", err)
		return
	}

	fmt.Printf("---------🔥🔥🔥----------\n")
	fmt.Printf("json after UnMarshall : \n %v \n", users)

	for k, v := range users {
		// fmt.Printf("%d : name => %s, telephone => %s, age => %d\n", k, v.Name, v.Phone, v.Age)
		fmt.Printf("~~~~~> user[%d]: %+v\n", k, v)
		fmt.Println("")
	}

	//-------------------------------------------------------

	var myStruct []User

	var user_one User
	user_one.Name = "Kevin"
	user_one.Age = 30
	user_one.Email = "kev@gmail.com"
	user_one.Phone = "0617283940"
	user_one.PostalCode = "75009"

	// on add les donnees de user a mystruct qui est de type User
	myStruct = append(myStruct, user_one)

	jsonFromSlice, err := json.MarshalIndent(myStruct, "", " ")

	fmt.Printf("---------🔥🔥🔥----------\n")
	fmt.Printf("------Json Marhsall-------\n")

	if err != nil {
		fmt.Println("Error Marshalling :", err)
	}
	fmt.Println(string(jsonFromSlice))

}
