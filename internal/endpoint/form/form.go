package form

import (
	"coven/internal/endpoint"
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
	}
}
