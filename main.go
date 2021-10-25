package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/yusuftatli/yemek-sepeti-api/handlers"
)



var keyValueStore *handlers.KeyValueStore

func main(){
	keyValueStore = handlers.NewKeyValuStore()


	ticker := time.NewTicker(time.Second * 10)
	go func() {
    
   for range ticker.C {
    handlers.WriteData(keyValueStore.GetAll())
   }

}()
	InitializeRoutes()

	fmt.Printf("Starting server at port 8000\n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
			log.Fatal(err)
	}

	
	
}



 


func writeFile(){

	 obj := map[string]string{}
	

 	obj["abc"]="def"
 	obj["abc2"]="def4"
 	obj["abc3"]="def5"

 	exampleBytes, err := json.Marshal(obj)
    if err != nil {
        log.Fatal(err)
    }

	err = ioutil.WriteFile("./db.text", exampleBytes, 0777)
	// handle this error
	if err != nil {
  // print it out
  fmt.Println(err)


		}
	}

	func readText(){
		 file, err := os.Open("./db.text")
		 if err != nil {
		     log.Fatal(err)
		 }
		 defer file.Close()
		  data := map[string]string{}
		 if err := json.NewDecoder(file).Decode(&data); err != nil {
		     log.Fatal(err)
		 }
		
	}

