package main

import (
	"hackaton-facef-api/api"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile) //Informa o local do erro
	api := api.App{}
	api.StartServer()
}
