package main

import (
	"anonichat/api"
	"anonichat/models"
)

func init() {
	models.InitDB()
}

func main() {
	r := api.InitRoutes()

	r.Run(":4545")

}
