package api

import (
	"encoding/json"
	"net/http"

	"github.com/svasandani/terrakita/internal/benchmark"
	"github.com/svasandani/terrakita/internal/db"
)

func FilterLingsHandler(w http.ResponseWriter, r *http.Request) {
	var f db.FilterLingsRequest
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

	benchmark.Start("FilterLings")
	fr, err := db.FilterLings(f)
	benchmark.Stop("FilterLings")
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

func FilterLingPropertiesHandler(w http.ResponseWriter, r *http.Request) {
	var f db.FilterLingPropertiesRequest
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

	benchmark.Start("FilterLingProperties")
	fr, err := db.FilterLingProperties(f)
	benchmark.Stop("FilterLingProperties")
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

func FilterLingletsHandler(w http.ResponseWriter, r *http.Request) {
	var f db.FilterLingletsRequest
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

	benchmark.Start("FilterLinglets")
	fr, err := db.FilterLinglets(f)
	benchmark.Stop("FilterLinglets")
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

func FilterLingletPropertiesHandler(w http.ResponseWriter, r *http.Request) {
	var f db.FilterLingletPropertiesRequest
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

	benchmark.Start("FilterLingletProperties")
	fr, err := db.FilterLingletProperties(f)
	benchmark.Stop("FilterLingletProperties")
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

func CompareLingsHandler(w http.ResponseWriter, r *http.Request) {
	var c db.CompareLingsRequest
	var js []byte
	var er error

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		js, er = errorResponse(err, http.StatusBadRequest)
		if er != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		writeResponse(w, js)
		return
	}

	benchmark.Start("CompareLings")
	cr, err := db.CompareLings(c)
	benchmark.Stop("CompareLings")
	if err != nil {
		js, er = errorResponse(err, http.StatusInternalServerError)
		if er != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		writeResponse(w, js)
		return
	}

	js, err = json.Marshal(cr)
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

func CompareLingletsHandler(w http.ResponseWriter, r *http.Request) {
	var c db.CompareLingletsRequest
	var js []byte
	var er error

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		js, er = errorResponse(err, http.StatusBadRequest)
		if er != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		writeResponse(w, js)
		return
	}

	benchmark.Start("CompareLinglets")
	cr, err := db.CompareLinglets(c)
	benchmark.Stop("CompareLinglets")
	if err != nil {
		js, er = errorResponse(err, http.StatusInternalServerError)
		if er != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		writeResponse(w, js)
		return
	}

	js, err = json.Marshal(cr)
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
		Message:    e.Error(),
		StatusCode: i,
	}

	js, err := json.Marshal(er)
	if err != nil {
		return nil, err
	}

	return js, nil
}
