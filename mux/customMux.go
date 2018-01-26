package customMux

import (
	"fmt"
	"math/rand"
	"net/http"
)

// CustomServerMux - Multiplexer structure
type CustomServerMux struct {
}

// Any object of type CustomServerMux will be able to call this function.
// ServeHTTP is actually set up in the Handler interface in server.go.
// Our CustomServerMux is going to overwrite ServeHTTP function.
func (p *CustomServerMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sendRandomNumber(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sendRandomNumber(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Here's a random number: %d", rand.Int())
}

// InitMux - Initialize our custom server multiplexer
func InitMux(addr string) {
	mux := &CustomServerMux{}
	http.ListenAndServe(addr, mux)
}
