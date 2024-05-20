package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/FerMusicComposer/totalcoder-challenge1/utils"
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
		fmt.Println(utils.NotAllowed.From("HandlePostData => method not allowed", nil).Err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	req := make(map[string]string)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(utils.BadRequest.From("HandlePostData => error reading request body", err).Err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, &req); err != nil {
		fmt.Println(utils.BadRequest.From("HandlePostData => error parsing request body", err).Err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	key, okKey := req["key"]
	value, okValue := req["value"]
	if !okKey || !okValue {
		fmt.Println(utils.InvalidPayload.From("HandlePostData => invalid payload", nil).Err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	mh.store[key] = value
	fmt.Println("Data saved:", mh.store)

	jsonResp, err := json.Marshal(req)
	if err != nil {
		fmt.Println(utils.Internal.From("HandleGetData => error marshalling data", err).Err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func (mh *InMemoryHandler) HandleGetData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Println(utils.NotAllowed.From("HandleGetData => method not allowed", nil).Err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	key := r.URL.Query().Get("key")
	fmt.Println("Key received:", key)
	value, ok := mh.store[key]
	if !ok {
		fmt.Println(utils.NotFound.From("HandleGetData => key not found", nil).Err)
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}

	resp := map[string]string{
		"key":   key,
		"value": value,
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(utils.Internal.From("HandleGetData => error marshalling data", err).Err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
