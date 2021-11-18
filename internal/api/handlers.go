package api

import (
	"encoding/json"
	"net/http"

	"github.com/svasandani/terrakita/internal/db"
)

func PostSearchFilterHandler(w http.ResponseWriter, r *http.Request) {
	var f db.FilterRequest

	err := json.NewDecoder(r.Body).Decode(&f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fr := db.FilterResponse{}
	js, err := json.Marshal(fr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
