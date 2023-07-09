package storage

import (
	"fmt"
	"key-value-storage/param"
	"key-value-storage/storage"
	"sync"
)

type Service struct {
}

var l = sync.RWMutex{}

func New() Service {
	return Service{}
}

func (s Service) Put(req param.PutInStoreRequest) param.PutInStorageResponse {
	l.Lock()
	_, ok := storage.Store[req.Key]
	storage.Store[req.Key] = req.Value
	l.Unlock()

	if ok {
		return param.PutInStorageResponse{Status: param.StatusUpdated}
	}

	return param.PutInStorageResponse{Status: param.StatusCreated}

}

func (s Service) Get(req param.GetFromStorageRequest) (param.GetFromStorageResponse, error) {
	l.RLock()
	if val, ok := storage.Store[req.Key]; ok {
		return param.GetFromStorageResponse{Value: val}, nil
	}

	return param.GetFromStorageResponse{}, fmt.Errorf("")
}

func (s Service) Delete(req param.DeleteFromStorageRequest) (param.DeleteFromStorageResponse, error) {

	if _, ok := storage.Store[req.Key]; ok {
		delete(storage.Store, req.Key)

		return param.DeleteFromStorageResponse{}, nil
	}

	return param.DeleteFromStorageResponse{}, fmt.Errorf("")
}
