package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// var keyValueStore *handlers.KeyValueStore

type FileInterface interface {
	WriteData() 
}

func WriteData(values map[string]string){

	exampleBytes, err := json.Marshal(values)
	 if err != nil {
			 log.Fatal(err)
	 }

 err = ioutil.WriteFile("../db.txt", exampleBytes, 0644)
 // handle this error
 if err != nil {
 // print it out
 fmt.Println(err)


	 }
	 fmt.Println("veri yazıldı")
}

func WriteDataToCache(){
	file, err := os.Open("../db.txt")
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close()
	var data []map[string]string
	if err := json.NewDecoder(file).Decode(&data); err != nil {
			log.Fatal(err)
	}
	// SetAll(data)
}

// t := time.NewTicker(time.Second * 1)
// go func() {
// 	for range t.C {
// 		c.reap()
// 	}
// }()

// func () reap() {
// 	t.lock.Lock()
// 	defer t.lock.Unlock()
// 	for k, v := range t.data {
// 		elapsed := time.Since(v.updated)
// 		if elapsed >= t.ttl {
// 			t.logger.Info(fmt.Sprintf("reaping key:%s", k))
// 			delete(t.data, k)
// 		}
// 	}
// }


func CreateFile(){
	f, e := os.Create("filename.txt")
	if e != nil {
		 panic(e)
	}
	defer f.Close()
	fmt.Fprintln(f, "Hello")
}