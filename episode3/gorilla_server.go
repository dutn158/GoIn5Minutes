package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type ServerStatus struct {
	NumReservations int  `json:"total_num_reservations"`
	Reserved        bool `json:"currently_reserved"`
}

var servers = map[string]*ServerStatus{
	"alice": &ServerStatus{},
	"bob":   &ServerStatus{},
	"carol": &ServerStatus{},
}

var mx = sync.RWMutex

func main() {
	router := mux.NewRouter()

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("%s not found\n", r.URL)))
	})

	apiRouter := router.Headers("Content-Type", "application/json").Subrouter()
	apiRouter.HandleFunc("/api/{name}", getServer).Methods("GET")
	apiRouter.HandleFunc("/api/{name}", reserveServer).Methods("POST")
	apiRouter.HandleFunc("/api/{name}", releaseServer).Methods("DELETE")

	adminAPIRouter := router.Headers("Content-Type", "application/json").MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
		adminToken := r.Header.Get("X-ADMIN-TOKEN")
		if adminToken == "" {
			return false
		}
		return adminToken == "SuperAdmin"
	}).SubRouter()

	adminAPIRouter.HandleFunc("/api/admin/servers", getAllServers).Methods("GET")
	adminAPIRouter.HandleFunc("/api/admin/servers", releaseAllServers).Methods("DELETE")

	log.Printf("serving on port 8080")
	http.ListenAndServe(":8080", router)
}
