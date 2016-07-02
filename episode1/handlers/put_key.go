package handlers

import (
	"net/http"

	"github.com/dutn158/GoIn5Minutes/episode1/storage"
)

// PutKey returns an http.Handler that can set a value for the key registered by Gorilla
// mux as "key" in the path. It expects the value to be in the body of the PUT request
func PutKey(db storage.DB) http.Handler {
	return http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")

		if key == "" {
			http.Error(w, Error("missing key name in path"), http.StatusBadRequest)
			return
		}

		defer r.Body.Close()

		if err != nil {
			http.Error(w, Error("error reading PUT body"), http.StatusBadRequest)
			return
		}

		if err := db.Set(key, val); err != nil {
			http.Error(w, Error("error setting value in DB"), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

	})
}
