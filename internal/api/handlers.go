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

func CrossLingPropertiesHandler(w http.ResponseWriter, r *http.Request) {
	var c db.CrossLingPropertiesRequest
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

	benchmark.Start("CrossLingProperties")
	cr, err := db.CrossLingProperties(c)
	benchmark.Stop("CrossLingProperties")
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

func CrossLingletPropertiesHandler(w http.ResponseWriter, r *http.Request) {
	var c db.CrossLingletPropertiesRequest
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

	benchmark.Start("CrossLingletProperties")
	cr, err := db.CrossLingletProperties(c)
	benchmark.Stop("CrossLingletProperties")
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

func ImplicationAntecedentHandler(w http.ResponseWriter, r *http.Request) {
	var i db.ImplicationRequest
	var js []byte
	var er error

	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		js, er = errorResponse(err, http.StatusBadRequest)
		if er != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		writeResponse(w, js)
		return
	}

	benchmark.Start("ImplicationAntecedent")
	ir, err := db.ImplicationAntecedent(i)
	benchmark.Stop("ImplicationAntecedent")
	if err != nil {
		js, er = errorResponse(err, http.StatusInternalServerError)
		if er != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		writeResponse(w, js)
		return
	}

	js, err = json.Marshal(ir)
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

func ImplicationConsequentHandler(w http.ResponseWriter, r *http.Request) {
	var i db.ImplicationRequest
	var js []byte
	var er error

	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		js, er = errorResponse(err, http.StatusBadRequest)
		if er != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		writeResponse(w, js)
		return
	}

	benchmark.Start("ImplicationConsequent")
	ir, err := db.ImplicationConsequent(i)
	benchmark.Stop("ImplicationConsequent")
	if err != nil {
		js, er = errorResponse(err, http.StatusInternalServerError)
		if er != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		writeResponse(w, js)
		return
	}

	js, err = json.Marshal(ir)
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

func ImplicationDoubleHandler(w http.ResponseWriter, r *http.Request) {
	var i db.ImplicationRequest
	var js []byte
	var er error

	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		js, er = errorResponse(err, http.StatusBadRequest)
		if er != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		writeResponse(w, js)
		return
	}

	benchmark.Start("ImplicationDouble")
	ir, err := db.ImplicationDouble(i)
	benchmark.Stop("ImplicationDouble")
	if err != nil {
		js, er = errorResponse(err, http.StatusInternalServerError)
		if er != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		writeResponse(w, js)
		return
	}

	js, err = json.Marshal(ir)
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

func SimilarityLingsHandler(w http.ResponseWriter, r *http.Request) {
	var s db.SimilarityLingsRequest
	var js []byte
	var er error

	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		js, er = errorResponse(err, http.StatusBadRequest)
		if er != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		writeResponse(w, js)
		return
	}

	benchmark.Start("SimilarityLings")
	sr, err := db.SimilarityLings(s)
	benchmark.Stop("SimilarityLings")
	if err != nil {
		js, er = errorResponse(err, http.StatusInternalServerError)
		if er != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		writeResponse(w, js)
		return
	}

	js, err = json.Marshal(sr)
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

func SimilarityLingletsHandler(w http.ResponseWriter, r *http.Request) {
	var s db.SimilarityLingletsRequest
	var js []byte
	var er error

	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		js, er = errorResponse(err, http.StatusBadRequest)
		if er != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		writeResponse(w, js)
		return
	}

	benchmark.Start("SimilarityLinglets")
	sr, err := db.SimilarityLinglets(s)
	benchmark.Stop("SimilarityLinglets")
	if err != nil {
		js, er = errorResponse(err, http.StatusInternalServerError)
		if er != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		writeResponse(w, js)
		return
	}

	js, err = json.Marshal(sr)
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
