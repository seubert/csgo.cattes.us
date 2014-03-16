package main

import (
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/ell/csgo.cattes.us/oauth2"
	"github.com/martini-contrib/render"
	"net/http"
)

func Index(r render.Render, params martini.Params, tokens oauth2.Tokens) {
	r.HTML(200, "index", nil)
}

func TestAuth(tokens oauth2.Tokens) string {
	if tokens.IsExpired() {
		return "not logged in"
	}

	return "logged in"
}

func LoggedIn(tokens oauth2.Tokens, w http.ResponseWriter, r *http.Request) {
	if !tokens.IsExpired() {
		profile, err := GetProfile(tokens.Access())

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(profile.Active)
	}

	http.Redirect(w, r, "/", 302)
}

func GetSongs() {
}

func UploadSongs() {
}

func GetMaps() {
}

func UploadMaps() {
}
