package tests

import (
	"testing"

	"github.com/yusuftatli/yemek-sepeti-api/handlers"
)

var keyValueStore *handlers.KeyValueStore

func TestSetGet(t *testing.T) {
	keyValueStore = handlers.NewKeyValuStore()
	keyValueStore.Set("url", "www.yemeksepeti.com.tr")

	result := keyValueStore.Get("url")

	if result != "www.yemeksepeti.com.tr" {
		t.Errorf("value not equel")
	}
}