package middleware

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

// VerifyContentType - Check requests against their MIME type.
// If the requests are not of type JSON, we prevent the
// request from going any further.
func VerifyContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-type") != "application/json" {
			log.Println("The request's content type: ", r.Header.Get("Content-type"))
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("415 - Unsupported Media Type"))
			return
		}
		handler.ServeHTTP(w, r)
	})
}

// SetServerTime - Add server timestamp for response cookie.
// Cookie is set for every response
func SetServerTime(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
		cookie := http.Cookie{
			Name:  "Server-Time(UTC)",
			Value: strconv.FormatInt(time.Now().Unix(), 10),
		}
		http.SetCookie(w, &cookie)
		log.Println("In server time cookie middleware")
	})
}
