package cloudinary

import (
	"flag"
	"log"
	"os"
	"sync"
)

type singleton struct {
}

var instance *Service
var once sync.Once

var accountKey = flag.String("schemaRegistryAPI", os.Getenv("Cloudinary_Account_Key"), "The schema registry API")
var secretKey = flag.String("schemaRegistryAPI", os.Getenv("Cloudinary_Secret_Key"), "The schema registry API")
var cloudName = flag.String("schemaRegistryAPI", os.Getenv("Cloudinary_Cloud_Name"), "The schema registry API")

func GetService() *Service {
	once.Do(func() {
		var err error
		instance, err = Dial("cloudinary://" + accountKey + ":" + secretKey + "@" + cloudName)
		if err != nil {
			log.Fatal(err)
		}
	})
	return instance
}
