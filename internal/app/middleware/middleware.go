package middleware

import "net/http"

type FuncMiddleware interface {
	AddFunc(*http.ServeMux) (http.HandlerFunc, error)
}

type Middlware interface {
	Add(*http.ServeMux) (http.Handler, error)
}
