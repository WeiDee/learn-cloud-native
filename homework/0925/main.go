package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type logWriter struct {
	http.ResponseWriter
	*http.Request
	statusCode int
}

func (lg *logWriter) writeHeader(code int) {
	lg.statusCode = code
	lg.ResponseWriter.WriteHeader(code)
}

func (lg *logWriter) writeLog() {
	fmt.Fprintf(os.Stdout, "Client IP is: %s\n", lg.Request.RemoteAddr)
	fmt.Fprintf(os.Stdout, "Respons code is: %d\n", lg.statusCode)
}

func newLogWriter(w http.ResponseWriter, req *http.Request) *logWriter {
	return &logWriter{w, req, http.StatusOK}
}

func main() {
	//resp with req header and log
	respWithHeader := func(w http.ResponseWriter, req *http.Request) {
		httpHeader := req.Header

		for k, v := range httpHeader {
			w.Header().Set(k, strings.Join(v, ","))
		}

		lg := newLogWriter(w, req)
		lg.writeLog()

		io.WriteString(w, "This is resp with request header!\n")
	}

	//resp ENV version
	version := func(w http.ResponseWriter, req *http.Request) {

		Key := "VERSION"

		v := os.Getenv(Key)

		w.Header().Set(Key, v)

		io.WriteString(w, "This is resp with ENV VERSION!\n")
	}

	//resp 200
	healthz := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "ok")
	}

	http.HandleFunc("/", respWithHeader)
	http.HandleFunc("/version", version)
	http.HandleFunc("/healthz", healthz)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
