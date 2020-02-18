package bootstrap

import (
	"fmt"
	"log"
	"github.com/joho/godotenv"
	"github.com/RomaBiliak/go_api_admin_blog/api/controllers"

)

var server = controllers.Server{}

func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.Initialize()


	server.Run(":3000")

}