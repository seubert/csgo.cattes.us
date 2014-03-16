package main

import (
	"encoding/json"
	"net/http"
	"crypto/tls"
	"io/ioutil"
	"errors"
)


type Profile struct {
	Username string
	SomethingAwful *SomethingAwful
	Steam *SteamAccount
}

type SomethingAwful struct {
	Username string `json:"username"`
	Url string `json:"url"`
	UserID string `json:"userid"`
	PostCount int `json:"postcount"`
	RegDate string `json:"regdate"`
}

type SteamAccount struct {
	Username string `json:"username"`
	Url string `json:"url"`
	Userid string `json:"userid"`
}

type User struct {
	Username string `json:"username"`
}


func Steam64ToSteamID(steam64 int64) string {
	return "todo"
}

func GetProfile(token string) (*Profile, error) {
	profile := new(Profile)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	var url = "https://somethingauthful.com/api/user/?access_token=" + token
	req, _ := http.NewRequest("GET", url, nil)
	var resp, err = client.Do(req)

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	var objmap map[string]*json.RawMessage
	err = json.Unmarshal(data, &objmap)

	var active bool
	err = json.Unmarshal(*objmap["active"], &active)

	if !active {
		err = errors.New("GoonAuth account is unverified")
		return profile, err
	}

	user := &User{}
	err = json.Unmarshal(*objmap["user"], &user)

	steam := &SteamAccount{}
	err = json.Unmarshal(*objmap["steam"], &steam)

	somethingawful := &SomethingAwful{}
	err = json.Unmarshal(*objmap["somethingawful"], &somethingawful)

	profile.Username = user.Username
	profile.Steam = steam
	profile.SomethingAwful = somethingawful

	return profile, err
}
