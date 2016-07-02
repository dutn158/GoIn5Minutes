package handlers

import "net/http"

func GetKey(db DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "missing key name in query string", http.StatusBadRequest)
			return
		}
	})
}
