package main

import (
	"github.com/vitorfranklin/go_alura-gin-api-rest/database"
	"github.com/vitorfranklin/go_alura-gin-api-rest/routes"
)

func main() {

	database.ConectaComBancoDeDados()

	routes.HandleRequest()
}
