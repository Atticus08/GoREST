package serveFiles

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// RunFileServer - Initialize and start the static file server
func RunFileServer(addr string) {
	router := httprouter.New()
	router.ServeFiles("/static/*filepath", http.Dir("/usr/local/static"))
	log.Fatal(http.ListenAndServe(addr, router))
}
