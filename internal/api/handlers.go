package api

import (
	"encoding/json"
	"net/http"

	"github.com/svasandani/terrakita/internal/db"
)

func PostSearchFilterHandler(w http.ResponseWriter, r *http.Request) {
  var f db.FilterRequest
  var js []byte
  var er error

  err := json.NewDecoder(r.Body).Decode(&f)
  if err != nil {
    js, er = errorResponse(err, http.StatusBadRequest)
    if er != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    writeResponse(w, js)
    return
  }

  fr, err := db.Filter(f)
  if err != nil {
    js, er = errorResponse(err, http.StatusInternalServerError)
    if er != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    writeResponse(w, js)
    return
  }

  js, err = json.Marshal(fr)
  if err != nil {
    js, er = errorResponse(err, http.StatusInternalServerError)
    if er != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    writeResponse(w, js)
    return
  }

  writeResponse(w, js)
}

func writeResponse(w http.ResponseWriter, js []byte) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(js)
}

func errorResponse(e error, i int) ([]byte, error) {
  er := db.ErrorResponse{
    Message: e.Error(),
    StatusCode: i,
  }
  
  js, err := json.Marshal(er)
  if err != nil {
    return nil, err
  }
  
  return js, nil
}
