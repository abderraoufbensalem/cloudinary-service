package cloudinary

import (
	"log"
	"sync"
)

type singleton struct {
}

var instance *Service
var once sync.Once

func GetService() *Service {
	once.Do(func() {
		var err error
		instance, err = Dial("cloudinary://your-account-key:your-secret-key@your-cloud-name")
		if err != nil {
			log.Fatal(err)
		}
	})
	return instance
}
