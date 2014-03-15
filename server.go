package main

import (
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/oauth2"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"log"
	"net/http"
)

type App struct {
	m       *martini.ClassicMartini
	address string
}

func GoonAuth(opts *oauth2.Options) martini.Handler {
	opts.AuthUrl = "https://somethingauthful.com/o/authorize/"
	opts.TokenUrl = "https://somethingauthful.com/o/token/"

	return oauth2.NewOAuth2Provider(opts)
}

func NewApp(address string) *App {
	app := new(App)
	app.address = address
	app.m = martini.Classic()

	app.SetupMiddleware()
	app.SetupRoutes()

	return app
}

func (app *App) RunServer() {
	fmt.Println("Running server on " + app.address)
	log.Fatal(http.ListenAndServe(app.address, app.m))
}

func (app *App) SetupMiddleware() {
	store := sessions.NewCookieStore([]byte("changeme123"))
	app.m.Use(sessions.Sessions("cattes_session", store))

	app.m.Use(GoonAuth(&oauth2.Options{
		ClientId:     "W1=p?2Cwc09Kzm-g-FopfK3;voqxxtvCoBoCdG4Z",
		ClientSecret: "aEG:k8tjvFhQL;KTevoPTI:DvkPtsGEn9q4V_y=MrR;@:Z0XqFKX_5xdzu8nhThG5mp8ibKf9-Uu9kznez@oVLJxIBCjZrWKj=9.8LaOgJKQHnUpKx-;TLwVSMzbPTK1",
		RedirectURL:  "http://localhost:3000/oauth2callback",
		Scopes:       []string{"read"},
	}))

	app.m.Use(render.Renderer())
	app.m.Use(martini.Static("static"))
}

func (app *App) SetupRoutes() {
	app.m.Get("/", Index)
	app.m.Get("/testauth", TestAuth)
}

func main() {
	app := NewApp("127.0.0.1:3000")
	app.RunServer()
}
