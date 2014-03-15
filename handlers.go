package main

import (
	"github.com/martini-contrib/render"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/oauth2"
)

func Index(r render.Render, params martini.Params) {
	r.HTML(200, "index", nil)
}

func TestAuth(tokens oauth2.Tokens) string {
	if tokens.IsExpired() {
		return "not logged in"
	}

	return "logged in"
}

func GetSongs() {
}

func UploadSongs() {
}

func GetMaps() {
}

func UploadMaps() {
}
