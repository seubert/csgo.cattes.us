package main

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Profile struct {
	User User
	SomethingAwful SomethingAwful
	Steam SteamAccount
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

type User struct {
	Username string
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

	r := &Profile{}
	err = json.Unmarshal(data, &r)

	profile.User = r.User
	profile.Steam = r.Steam
	profile.SomethingAwful = r.SomethingAwful

	return profile, err
}
