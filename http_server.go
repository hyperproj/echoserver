package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func handleHTTP(w http.ResponseWriter, r *http.Request) {
	hostName, _ := os.Hostname()
	printf(w, "Hostname: %s\n", hostName)

	printf(w, "\nRequest Info:\n")
	printf(w, "    TLS: %t\n", r.TLS != nil)
	printf(w, "    HTTP Version: %s\n", r.Proto)
	printf(w, "    HTTP Method: %s\n", r.Method)
	printf(w, "    HTTP URI: %s\n", r.RequestURI)
	printf(w, "    HTTP URL Path: %s\n", r.URL.Path)
	printf(w, "    HTTP Query: %s\n", r.URL.RawQuery)
	printf(w, "    Content-Length: %d\n", r.ContentLength)
	printf(w, "    Downstream IP: %s\n", getDownstreamIP(r))

	printf(w, "\n  Headers:\n")
	for key, values := range r.Header {
		printf(w, "    %s: %s\n", key, values)
	}

	printf(w, "\n  Body:\n")
	if r.Body != nil && r.ContentLength > 0 {
		defer r.Body.Close()
		bs, err := io.ReadAll(r.Body)

		if err != nil {
			printf(w, "    Read body error: %s\n", err.Error())
		} else {
			printf(w, "    %s", bs)
		}
	} else {
		printf(w, "    No Body\n")
	}

	printf(w, "\nEnvironment Variables:\n")
	for _, env := range os.Environ() {
		printf(w, "    %s\n", env)
	}

	printf(w, "\n")
}

func StartHTTPServer(port int) {
	http.HandleFunc("/", handleHTTP)
	addr := fmt.Sprintf(":%d", port)
	server := &http.Server{
		Addr:    addr,
		Handler: http.HandlerFunc(handleHTTP),
	}

	fmt.Printf("Starting http1.0, http1.1 echoserver on %s...\n", addr)
	server.ListenAndServe()
}
