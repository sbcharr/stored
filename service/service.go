package service

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	kv "github.com/sbcharr/stored/pkg/persistence/memory"
)

var productDB = kv.NewKVStore()

// KVPutHandler handles all upsert actions to the KV store
func KVPutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = productDB.Put(key, string(value))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// KVGetHandler handles all requests to get value from the KV store
func KVGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := productDB.Get(key)
	if err != nil {
		if errors.Is(err, kv.ErrKeyNotFound) {
			http.Error(w, kv.ErrKeyNotFound.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	w.Write([]byte(value))
}

// KVDeleteHandler handles all delete operations to the KV store
func KVDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	err := productDB.Delete(key)
	if err != nil {
		if errors.Is(err, kv.ErrKeyNotFound) {
			http.Error(w, kv.ErrKeyNotFound.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	w.WriteHeader(http.StatusOK)
}
