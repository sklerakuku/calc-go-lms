package main

import (
	"github.com/sklerakuku/calc-go-lms/internal/application"
)

func main() {

	app := application.New()
	//app.Run()
	app.RunServer()
}
