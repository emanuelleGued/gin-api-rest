package main

import (
	"github.com/emanuelleGued/api-go-gin/database"
	"github.com/emanuelleGued/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()

}
