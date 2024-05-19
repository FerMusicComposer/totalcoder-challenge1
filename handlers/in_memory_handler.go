package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type InMemoryHandler struct {
	store map[string]string
}

func NewInMemoryHandler(store map[string]string) *InMemoryHandler {
	return &InMemoryHandler{
		store: store,
	}
}

func (mh *InMemoryHandler) HandlePostData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	req := make(map[string]string)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	key, okKey := req["key"]
	value, okValue := req["value"]
	if !okKey || !okValue {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	mh.store[key] = value
	fmt.Println("Data saved:", mh.store)
	jsonResp, err := json.Marshal(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func (mh *InMemoryHandler) HandleGetData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	key := r.URL.Query().Get("key")
	fmt.Println("Key received:", key)
	value, ok := mh.store[key]
	if !ok {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}

	resp := map[string]string{
		"key":   key,
		"value": value,
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
