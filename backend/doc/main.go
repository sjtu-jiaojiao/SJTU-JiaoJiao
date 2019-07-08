package main

import (
	"jiaojiao/utils"
)

func main() {
	router, rg := utils.CreateAPIGroup()

	rg.Static("/doc", "./doc")

	utils.RunWebService("doc", router)
}
