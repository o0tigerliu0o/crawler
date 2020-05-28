package model

import "encoding/json"

type Profile struct {
	Name       string
	Gender     string
	Age        int
	Height     int
	Weight     string
	Income     string
	Marriage   string
	Education  string
	Occupation string
	WorkDest   string
	Hokou      string
	Xinzuo     string
	House      string
	Car        string
}

// 用于将interface转成profile
func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}

	err = json.Unmarshal(s, &profile)
	return profile, err
}
