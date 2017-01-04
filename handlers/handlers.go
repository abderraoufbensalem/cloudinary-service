package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ricardo-ch/platform-poc/app/sample-cloudinary/cloudinary"
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
