package main

import (
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/ell/csgo.cattes.us/oauth2"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"net/http"
)

func Index(r render.Render, params martini.Params, tokens oauth2.Tokens) {
	r.HTML(200, "index", nil)
}

func LoggedIn(tokens oauth2.Tokens, w http.ResponseWriter, r *http.Request, session sessions.Session) {
	if !tokens.IsExpired() {
		profile, err := GetProfile(tokens.Access())

		if !profile.Active {
		}

		if err != nil {
			fmt.Println(err)
		}
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
