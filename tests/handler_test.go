package tests

import (
	"testing"

	"github.com/yusuftatli/yemek-sepeti-api/handlers"
)

var keyValueStore *handlers.KeyValueStore
var fileHandler *handlers.FileHandler

func TestSetGet(t *testing.T) {
	keyValueStore = handlers.NewKeyValuStore()
	keyValueStore.Set("url", "www.yemeksepeti.com.tr")

	result := keyValueStore.Get("url")

	if result != "www.yemeksepeti.com.tr" {
		t.Errorf("value set incorrect")
	}
}

func TestInitilizeCache(t *testing.T) {
	var obj = make(map[string]string)
	 obj["value1"] = "value1"
	 obj["value2"] = "value2"

	 keyValueStore.InitilizeCache(obj)
	 allData := keyValueStore.GetAll()

	 if len(obj) !=len(allData) {
	   t.Errorf("initilize cache doesnt work")
	 }
}

 func TestFileLoad(t *testing.T) {
	keyValueStore.Set("url", "www.yemeksepeti.com.tr")
 	fileHandler.WritePeriodically()

	fileHandler.LoadData()
	result := keyValueStore.Get("url")

	if result != "www.yemeksepeti.com.tr" {
		t.Errorf("file load doesnt run")
	}
 }