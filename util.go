package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
)

func getDownstreamIP(r *http.Request) string {
	clientIP, _, _ := net.SplitHostPort(r.RemoteAddr)
	return clientIP
}

func printf(w io.Writer, format string, a ...any) {
	fmt.Printf(format, a...)
	fmt.Fprintf(w, format, a...)
}
