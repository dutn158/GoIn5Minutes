package handlers

import (
	"net/http"

	"github.com/dutn158/GoIn5Minutes/episode1/storage"
)

func PutKey(db storage.DB) http.Handler {
	return http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")

		if key == "" {
			http.Error(w, "missing key name in path", http.StatusBadRequest)
			return
		}

		defer r.Body.Close()

		if err != nil {
			http.Error(w, "error reading PUT body", http.StatusBadRequest)
			return
		}

		if err := db.Set(key, val); err != nil {
			http.Error(w, "error setting value in DB", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

	})
}
