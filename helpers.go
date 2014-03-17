package main

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Profile struct {
	User           struct{ Username string }
	SomethingAwful SomethingAwful
	Steam          SteamAccount
	Active         bool
}

type SomethingAwful struct {
	Username  string
	Url       string
	UserID    string
	PostCount int
	RegDate   string
}

type SteamAccount struct {
	Username string
	Url      string
	Userid   string
}

func (profile *Profile) ToJson() (string, error) {
	data, err := json.Marshal(profile)
	return string(data[:]), err
}

func (profile *Profile) FromJson(jsonData string) (*Profile, error) {
	err := json.Unmarshal([]byte(jsonData), &profile)
	return profile, err
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
	err = json.Unmarshal(data, &profile)

	return profile, err
}
