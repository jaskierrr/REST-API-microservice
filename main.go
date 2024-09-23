package main

import (
	"card-project/bootstrapper"
	"log"
)

func main() {
	err := bootstrapper.New().RunAPI()
	if err != nil {
		log.Fatal("failed to start")
	}
}
