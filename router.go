package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/yusuftatli/yemek-sepeti-api/handlers"
)

//todo:request reponse log
//todo corelation id headeran alınıp loga yazılacak

func InitializeRoutes(keyValueStore *handlers.KeyValueStore){

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		params := strings.Split(r.URL.RawQuery,"=")

		log.Println("co id =>",GetCorrelation(r))
		if len(params) != 2 {
			bin, _ := json.Marshal("Url Params 'key' is not valid")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write(bin)
		return
		}

    result :=	keyValueStore.Get(params[1])
		bin, _ := json.Marshal(result)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(bin)
	}) 

	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		params := strings.Split(r.URL.RawQuery,"=")
		log.Println("co id =>",GetCorrelation(r))

		if len(params) != 2 {
			bin, _ := json.Marshal("Url Params 'key' is not valid")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write(bin)
			return
		}

  	keyValueStore.Set(params[0],params[1])
		bin, _ := json.Marshal("success")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(bin)
	}) 

	http.HandleFunc("/deleteall", func(w http.ResponseWriter, r *http.Request) {
		log.Println("co id =>",GetCorrelation(r))
		keyValueStore.DeleteAllCache()
		bin, _ := json.Marshal("deleted all cache values")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(bin)
	}) 
}

func GetCorrelation(r *http.Request)string{
	id := r.Header.Get("X-Correlation-Id")
		if id == "" {
			newid := uuid.New()
			id = newid.String()
		}
		return id
}