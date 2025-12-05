package endpoint

import "net/http"

var Scheme string
var Address string
var Port uint16

type Endpoint struct {
	Path        string
	Methods     []string
	Secure      bool
	HandlerFunc http.HandlerFunc
}
