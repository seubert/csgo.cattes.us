package main

import (
	"github.com/codegangsta/martini"
	"github.com/ell/csgo.cattes.us/oauth2"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"net/http"
)

func Index(r render.Render, params martini.Params, session sessions.Session) {
	profileJson, ok := session.Get("Profile").(string)
	profile := new(Profile)

	if ok {
		profile.FromJson(profileJson)
	}

	data := struct {
		Profile *Profile
	}{
		profile,
	}

	r.HTML(200, "index", data)
}

func LoggedIn(tokens oauth2.Tokens, w http.ResponseWriter, r *http.Request, session sessions.Session) {
	if !tokens.IsExpired() {
		profile, err := GetProfile(tokens.Access())
		profileJson, err := profile.ToJson()

		if err != nil {
			session.AddFlash(err)
			session.Delete("oauth2_token")
		} else if !profile.Active {
			session.AddFlash("Please verify your GoonAuth account before logging in.")
			session.Delete("oauth2_token")
		} else if profile.Steam.Userid == "None" {
			session.AddFlash("Please add a Steam Account to you GoonAuth profile before logging in.")
			session.Delete("oauth2_token")
		} else {
			session.Set("Profile", profileJson)
			session.AddFlash("Successfully logged in as " + profile.User.Username)
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
