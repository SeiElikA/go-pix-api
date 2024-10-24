package main

import (
	"go-pix-api/src/config"
	"go-pix-api/src/routes"
	"go-pix-api/src/validation"
)

func main() {
	config.ConnectDB()
	validation.RegisterValidations()
	r := routes.SetupRouter()

	r.MaxMultipartMemory = 128 << 20
	r.Static("/profile_image", "./profile_image")
	r.Static("/post", "./post")

	r.Run(":8080")
}
