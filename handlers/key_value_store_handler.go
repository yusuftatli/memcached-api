package handlers

import (
	"sync"
)

//data local cache
//lock it locks the source while set process same time
type KeyValueStore struct {
	data   map[string]string
	lock   *sync.Mutex 
}

//create new instance
func NewKeyValuStore() (*KeyValueStore) {
	c := &KeyValueStore{
		data:   map[string]string{},
		lock:   &sync.Mutex{},
	}
	return c
}

//given key and value writes to the local cache
func (t *KeyValueStore) Set(key string, value string) error {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.data[key] = value
	return nil
}

//return specific cache value given by key value
func (t *KeyValueStore) Get(key string) string {
	if item, ok := t.data[key]; ok {
		return item
	}
	return ""
}

//returns all cache values
func (t *KeyValueStore) GetAll() map[string]string {
	return t.data
}

func (t *KeyValueStore) InitilizeCache(values map[string]string)  {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.data = values
}

func (t *KeyValueStore) DeleteAllCache()  {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.data = nil
}