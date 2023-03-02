package main

import (
	"github.com/MochammadQemalFirza/echo_web/routes"
)

func main() {
	e := routes.InitRouter()

	e.Logger.Fatal(e.Start(":8081"))
}
