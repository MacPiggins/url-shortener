package noDB

import (	
	//"fmt"
	"sync"
	"url-shortener/internal/database"
)

type NoDB struct {
	storage map[string]string
	mtx     *sync.RWMutex
}

func New() (*NoDB, error) {
	return &NoDB{
		storage: map[string]string{}, mtx: &sync.RWMutex{}}, nil
}

func (n *NoDB) Set(token, link string) error {	
	n.mtx.Lock()
	defer n.mtx.Unlock()

	if _, ok := n.storage[token]; ok {
		return database.UniqueError{}
	}

	n.storage[token] = link

	return nil
}

func (n *NoDB) Get(token string) (string, error) {
	n.mtx.RLock()
	defer n.mtx.RUnlock()

	if _, ok := n.storage[token]; !ok {
		return "", database.NotFoundError{}
	}	

	return n.storage[token], nil
}

func (n *NoDB) Close() {
	//do nothing
}
