package main

import (
	"database/sql"
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/ell/csgo.cattes.us/oauth2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"net/http"
)

func Index(r render.Render, params martini.Params, session sessions.Session) {
	profileJson, ok := session.Get("Profile").(string)
	profile := new(Profile)

	messages := session.Flashes()
	if len(messages) <= 0 {
		messages = nil
	}

	fmt.Println(messages)

	if ok {
		profile.FromJson(profileJson)
	}

	data := struct {
		Profile  *Profile
		Messages []interface{}
	}{
		profile,
		messages,
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

func GoonAuth(opts *oauth2.Options) martini.Handler {
	opts.AuthUrl = "https://somethingauthful.com/o/authorize/"
	opts.TokenUrl = "https://somethingauthful.com/o/token/"

	return oauth2.NewOAuth2Provider(opts)
}

func DB() martini.Handler {
	db, err := sql.Open("mysql", "root:penis123@tcp(127.0.0.1:3306)/csgo")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}

	return func(c martini.Context) {
		c.Map(db)
		c.Next()
	}
}
