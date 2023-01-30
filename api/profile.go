package api

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func (api *API) ImgProfileView(w http.ResponseWriter, r *http.Request) {
	// View with response image `img-avatar.png` from path `assets/images`
	// TODO: answer here
	bytePict, err := ioutil.ReadFile("./assets/images/img-avatar.png")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(bytePict)

}

func (api *API) ImgProfileUpdate(w http.ResponseWriter, r *http.Request) {
	// Update image `img-avatar.png` from path `assets/images`
	alias := r.FormValue("alias")

	uploadedFile, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer uploadedFile.Close()
	dir, err := os.Getwd()
if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
}

//membuat nama file baru yang akan disimpan
filename := handler.Filename
if alias != "" { //kalau alias disi maka nama file = alias
    filename = fmt.Sprintf("%s%s", alias, filepath.Ext(handler.Filename))
}

//membentuk lokasi tempat menyimpan file
fileLocation := filepath.Join(dir, "images", filename)
targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
}
defer targetFile.Close()

//mengisi file baru dengan data dari file yang ter-upload
if _, err := io.Copy(targetFile, uploadedFile); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
}

w.Write([]byte("done"))
	// TODO: answer here

	api.dashboardView(w, r)
}
