package main

import (
	"fmt"
	"html"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Atticus08/GoLangFun/GoREST/models"
)

func romanNumServer() {
	// Handle the HTTP request
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Split the URL into different strings using "/" as the delimiter
		urlElements := strings.Split(r.URL.Path, "/")
		if urlElements[1] == "roman_number" {
			// If the GET request contains "roman_numer" in its uri
			num, _ := strconv.Atoi(strings.TrimSpace(urlElements[2]))

			// Our roman numerals only go up to 5. If our number in the
			// url is greater than 5, then we need to send "Not Found" status
			if num == 0 || num > 5 {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - Number Not Found"))
			} else {
				fmt.Fprintf(w, "%q", html.EscapeString(romanNumerals.Numerals[num]))
			}
		} else {
			// All other request are BAD requests!
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad Request"))
		}
	})

	initServer()
}

func initServer() {
	server := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
