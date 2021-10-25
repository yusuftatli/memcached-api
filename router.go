package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// todo:request reponse log
//todo corelation id headeran alınıp loga yazılacak


func InitializeRoutes(){
 
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		params := strings.Split(r.URL.RawQuery,"=")

		if  len(params) !=2 {
			bin, _ := json.Marshal("Url Params 'key' is not valid")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write(bin)
		return
		}

    result :=	keyValueStore.Get(params[1])
		bin, _ := json.Marshal(result)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write(bin)
	}) 

	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		params := strings.Split(r.URL.RawQuery,"=")
		fmt.Println(params)
    
    fmt.Println("len=> ",len(params))

		if  len(params) !=2 {
			bin, _ := json.Marshal("Url Params 'key' is not valid")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			_, _ = w.Write(bin)
			return
		}

		ctx := r.Context()
		reqId := GetRequestID(ctx)
		log.Println(reqId)

  	 keyValueStore.Set(params[0],params[1])
		 bin, _ := json.Marshal("success")
		 w.Header().Set("Content-Type", "application/json")
		 w.WriteHeader(http.StatusCreated)
		 _, _ = w.Write(bin)
	}) 
	


	http.HandleFunc("/getall", func(w http.ResponseWriter, r *http.Request) {
		

  	data :=	 keyValueStore.GetAll()
		 bin, _ := json.Marshal(data)
		 w.Header().Set("Content-Type", "application/json")
		 w.WriteHeader(http.StatusCreated)
		 _, _ = w.Write(bin)
	}) 
}
type ContextKey string
const ContextKeyRequestID ContextKey = "requestID"
func GetRequestID(ctx context.Context) string {

	reqID := ctx.Value(ContextKeyRequestID)

	if ret, ok := reqID.(string); ok {
		return ret
	}

	return ""
}