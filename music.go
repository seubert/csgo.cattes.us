package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	id3 "github.com/mikkyang/id3-go"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type Response struct {
	Answer string `json:"answer"`
}

type Upload struct {
	Uploader string `json:"username"`
	Artist   string `json:"artist"`
	Genre    string `json:"genre"`
	FileName string `json:"file_name"`
	Title    string `json:"title"`
}

func GetMusic() {
}

func MusicUploader(r render.Render, params martini.Params) {
	r.HTML(200, "music_upload", nil)
}

func ParseMusicUpload(r *http.Request, session sessions.Session, db *sql.DB) (int, string) {
	err := r.ParseMultipartForm(100000)
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	files := r.MultipartForm.File["file"]

	for i, _ := range files {
		file, err := files[i].Open()
		defer file.Close()

		if err != nil {
			return http.StatusInternalServerError, err.Error()
		}

		songPath, _ := filepath.Abs("./public/uploads/music" + files[i].Filename)
		dst, err := os.Create(songPath)

		defer dst.Close()

		if err != nil {
			return http.StatusInternalServerError, err.Error()
		}

		if _, err := io.Copy(dst, file); err != nil {
			return http.StatusInternalServerError, err.Error()
		}

		profileJson := session.Get("Profile").(string)
		profile := new(Profile)
		profile.FromJson(profileJson)

		err = db.Ping()
		if err != nil {
			fmt.Println(err)
		}

		go HandleSongFile(songPath, files[i].Filename, profile, db)
	}

	response := Response{"File Transfer Completed"}
	j, err := json.Marshal(response)

	return 200, string(j[:])
}

func HandleSongFile(path, fileName string, profile *Profile, db *sql.DB) {
	mp3File, err := id3.Open(path)
	defer mp3File.Close()

	if err != nil {
		os.Remove(path)
		return
	}

	upload := new(Upload)
	upload.Uploader = profile.User.Username
	upload.FileName = fileName
	upload.Title = mp3File.Title()
	upload.Artist = mp3File.Artist()
	upload.Genre = mp3File.Genre()

	fmt.Println(upload.Title)

	row, _ := db.Query("INSERT INTO music (file_name, uploader, artist, title, genre) VALUES ($1, $2, $3, $4, $5)",
		upload.FileName,
		upload.Uploader,
		upload.Artist,
		upload.Title,
		upload.Genre,
	)

	defer row.Close()
	return
}
