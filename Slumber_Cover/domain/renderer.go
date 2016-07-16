package domain

import (
	"net/http"
)

type IRenderer interface {
	Render(w http.ResponseWriter, h *http.Request, status int, v interface{})
	JSON(w http.ResponseWriter, status int, v interface {})
	XML(w http.ResponseWriter, status int, v interface{})
	Data(w http.ResponseWriter, status int, v []byte)
	Text(w http.ResponseWriter, status int, v []byte)
}
