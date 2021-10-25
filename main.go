package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yusuftatli/yemek-sepeti-api/common"
	"github.com/yusuftatli/yemek-sepeti-api/handlers"
)

func main() {
	env := common.GetEnvironment()
	keyValueStore := handlers.NewKeyValuStore()
  fileHandler := handlers.NewFileHandler(keyValueStore)
	fileHandler.LoadData()
  fileHandler.WritePeriodically()
	InitializeRoutes(keyValueStore)

	fmt.Printf("Starting server at port " + env.ApplicationPort)
	if err := http.ListenAndServe(":" + env.ApplicationPort, nil); err != nil {
		log.Fatal(err)
	}
}