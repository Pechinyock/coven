package form

import (
	"coven/internal/endpoint"
	"path"
)

const FormPrefix = "/form"

func GetFormEndpoints() []endpoint.Endpoint {
	return []endpoint.Endpoint{
		{
			Path:        path.Join(FormPrefix, "card"),
			Methods:     []string{"POST", "GET", "PUT", "DELETE"},
			Secure:      true,
			HandlerFunc: nil,
		},
		{
			Path:        path.Join(FormPrefix, "image"),
			Methods:     []string{"POST", "GET", "DELETE"},
			Secure:      true,
			HandlerFunc: nil,
		},
	}
}
