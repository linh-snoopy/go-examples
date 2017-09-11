package main

import (
	"encoding/json"
	"log"
	"fmt"
)

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"first"` // want to change this to `json:"name"`
	tag  string `json:"-"`
	Another
}

type Another struct {
	Address string `json:"address"`
}

type MyUser struct {
	U User
}

func (u MyUser) MarshalJSON() ([]byte, error) {
	// encode the original
	m, _ := json.Marshal(u.U)

	// decode it back to get a map
	var a interface{}
	json.Unmarshal(m, &a)
	b := a.(map[string]interface{})

	// Replace the map key
	b[u.U.tag] = b["first"]
	delete(b, "first")

	// Return encoding of the map
	return json.Marshal(b)
}

func (u *MyUser) UnmarshalJSON(b []byte) error {
	// convert data to struct 
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	m["first"] = m[u.U.tag]
	delete(m, u.U.tag)
	
	p, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(p, &u.U)
}

type SearchResult struct {
	Date        *string      `json:"date,omitempty"`
	IdCompany   *int         `json:"idCompany,omitempty"`
	Company     *string      `json:"company,omitempty"`
	IdIndustry  *interface{} `json:"idIndustry,omitempty"`
	Industry    *string      `json:"industry,omitempty"`
	IdContinent *interface{} `json:"idContinent,omitempty"`
	Continent   *string      `json:"continent,omitempty"`
	IdCountry   *interface{} `json:"idCountry,omitempty"`
	Country     *string      `json:"country,omitempty"`
	IdState     *interface{} `json:"idState,omitempty"`
	State       *string      `json:"state,omitempty"`
	IdCity      *interface{} `json:"idCity,omitempty"`
	City        *string      `json:"city,omitempty"`
}

func main() {
	anoth := Another{"123 Jennings Street"}
	u := User{1, "Ken Jennings", "name", anoth}
	ju, _ := json.Marshal(u)
	jmu, _ := json.Marshal(MyUser{u})
	log.Printf("%s\n", ju)
	log.Printf("%s\n", jmu)
	j :=  `{"id":1,"name":"Ken Jennings","address":"456 Jennings Street"}`
	u2 := MyUser{
		User{tag: "name"},
	}
	json.Unmarshal([]byte(j), &u2)
	log.Println(u2.U)
	
	var result SearchResult
	date := "to be honest you should probably use a time.Time field here, just sayin"
	industry := "rocketships"
	var idCity interface{} = "interface{} is kinda inspecific, but this is the idcity field"
	city := "New York Fuckin' City"

	result.Date = &date
	result.Industry = &industry
	result.IdCity = &idCity
	result.City = &city

	b, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
}
