package main

import (
	"github.com/MochammadQemalFirza/echo_web/config"
	"github.com/MochammadQemalFirza/echo_web/routes"
)

func main() {
	config.CreateCon()
	e := routes.InitRouter()

	e.Logger.Fatal(e.Start(":8081"))
}
