package middleware

import "net/http"

type Middlware interface {
	Add(http.Handler) (http.Handler, error)
}
