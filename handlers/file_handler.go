package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/yusuftatli/yemek-sepeti-api/common"
)

type FileHandler struct {
	keyValueStore *KeyValueStore
}

type FileInterface interface {
	WritePeriodically()
	LoadData()
}

func NewFileHandler(keyValueStore *KeyValueStore) *FileHandler {
	return &FileHandler{keyValueStore}
}

//writes data from the local cache, given time value as a async
func (f *FileHandler) WritePeriodically() {
	timeInterval, err := strconv.Atoi(common.GetEnvironment().TimeInterval)
	if err != nil {
		timeInterval = 10
	}
	ticker := time.NewTicker(time.Duration(timeInterval) * time.Second)
	go func() {
		for range ticker.C {
			writeData(f.keyValueStore)
		}
	}()
}

//read data from the file and then write to local cache
func (f *FileHandler) LoadData() {
	if _, err := os.Stat("./db.txt"); os.IsNotExist(err) {
		file, err := os.Create("./db.txt")
		if err != nil {
			panic(err)
		}
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		var data map[string]string
		if err := json.NewDecoder(file).Decode(&data); err != nil {
			log.Fatal(err)
		}
		f.keyValueStore.InitilizeCache(data)
	} else {
		file, _ := os.OpenFile("./db.txt", os.O_RDWR|os.O_APPEND, 0660)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		var data map[string]string
		if err := json.NewDecoder(file).Decode(&data); err != nil {
			log.Fatal(err)
		}
		f.keyValueStore.InitilizeCache(data)
	}
}

//values takes from the local cache end writes to the local file
func writeData(keyValueStore *KeyValueStore) {
	exampleBytes, err := json.Marshal(keyValueStore.GetAll())
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(common.GetEnvironment().FilePath, exampleBytes, 0777)
	if err != nil {
		fmt.Println(err)
	}
}
