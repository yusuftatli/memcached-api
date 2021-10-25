package handlers

import (
	"sync"
)


type KeyValueStore struct {
	data   map[string]string //datanın tutulacağı local obje
	lock   *sync.Mutex //aynı zamanda set işlemi sırasında kaynağı locklar
}

//yeni instace oluşturur. 
func NewKeyValuStore() (*KeyValueStore) {
	c := &KeyValueStore{
		data:   map[string]string{},
		lock:   &sync.Mutex{},
	}
	return c
}

func (t *KeyValueStore) Set(key string, value string) error {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.data[key] = value
	return nil
}

func (t *KeyValueStore) Get(key string) string {
	if item, ok := t.data[key]; ok {
		return item
	}
	return ""
}

func (t *KeyValueStore) GetAll() map[string]string {
	return t.data
}

func (t *KeyValueStore) SetAll(values map[string]string)  {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.data = values
}