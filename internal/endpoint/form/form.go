package form

import (
	"coven/internal/endpoint"
	"net/http"
	"path"
)

const FormPrefix = "/form"

func GetFormEndpoints() []endpoint.Endpoint {
	return []endpoint.Endpoint{
		{
			Path:        "/card",
			Methods:     []string{"POST", "GET", "PUT", "DELETE"},
			Secure:      true,
			HandlerFunc: cardHandleFunc,
		},
		{
			Path:        path.Join(FormPrefix, "image"),
			Methods:     []string{"POST", "GET", "DELETE"},
			Secure:      true,
			HandlerFunc: imagePoolFileFunc,
		},
		{
			Path:        path.Join(FormPrefix, "push-changes"),
			Methods:     []string{"POST"},
			Secure:      true,
			HandlerFunc: pushChanges,
		},
		{
			Path:        "/partial",
			Methods:     []string{http.MethodPost, http.MethodGet, http.MethodDelete, http.MethodPatch},
			Secure:      true,
			HandlerFunc: HandlePartial,
		},
	}
}
