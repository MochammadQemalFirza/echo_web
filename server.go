package main

func main() {
	e := routes.init()

	e.Logger.Fatal(e.Start(":8081"))
}
