package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/abderraoufbensalem/cloudinary-service/cloudinary"
	"github.com/gorilla/mux"
)

type imageResource struct {
	publicId string `json:"public_id"`
}

func GetImageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//log.Println(r.Header)
	id := r.Header.Get("public_id")

	fmt.Printf("get image using public ID (%s)\n", id)

	url := cloudinary.GetService().Url(id, "t_450")

	j, _ := json.Marshal(url)
	w.Write(j)
}

func GetCatHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]
	fmt.Printf("get cat ID (%s)\n", id)

	fmt.Printf("get image using public ID (%s)\n", id)

	url := cloudinary.GetService().Url(id, "t_450")

	htmlResponse := fmt.Sprintf(`<!DOCTYPE html><html lang="en"><head></head><body><img src="%s"></body>`, url)

	w.Header().Set("Content-Type", "text/html")

	w.Write(htmlResponse)
}

func PostImageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	publicId, err := cloudinary.GetService().UploadFile("./testimages/ford4.jpg", nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Uploaded " + publicId)

	j, _ := json.Marshal(publicId)
	w.Write(j)
}
