package main

import (
	"app/crapp/router"
)

func main () {

	router := router.InitRouter()

	router.Run(":8000")
}