package domain

import (
	"net/http"
)

type ContextHandlerFunc func(http.ResponseWriter, *http.Request, IContext)

func (h ContextHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request, ctx IContext) {
	h(w, r, ctx)
}

type MiddlewareFunc func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)

func (m MiddlewareFunc) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	m(w, r, next)
}

type ContextMiddlewareFunc func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc, ctx IContext)

func (m ContextMiddlewareFunc) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc, ctx IContext) {
	m(w, r, next, ctx)
}

type IMiddleware interface {
	Handler(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
}

type IContextMiddleware interface {
	Handler(w http.ResponseWriter, r *http.Request, next http.HandlerFunc, ctx IContext)
}
