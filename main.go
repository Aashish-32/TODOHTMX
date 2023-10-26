package main

import (
	"github.com/Aashish32/htmx/model"
	"github.com/Aashish32/htmx/routes"
)

func main() {

	model.Setup()
	routes.SetupAndRun()

}
