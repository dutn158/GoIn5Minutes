package domain

import (
	"net/http"
)

type IResource interface {
	Context() IContext
	Routes() *Routes
	Render(w http.ResponseWriter, r *http.Request, status int, v interface{})
}


