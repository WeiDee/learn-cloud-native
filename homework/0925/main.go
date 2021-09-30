package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	//resp with req header
	respWithHeader := func(resp http.ResponseWriter, req *http.Request) {
		httpHeader := req.Header

		for k, v := range httpHeader {
			resp.Header().Set(k, strings.Join(v, ","))
		}

		io.WriteString(resp, "This is resp with request header!\n")
	}

	//resp ENV version
	version := func(resp http.ResponseWriter, req *http.Request) {

		Key := "VERSION"

		v := os.Getenv(Key)

		resp.Header().Set(Key, v)

		io.WriteString(resp, "This is resp with ENV VERSION!\n")
	}

	//resp 200
	healthz := func(resp http.ResponseWriter, req *http.Request) {
		resp.WriteHeader(http.StatusOK)
		io.WriteString(resp, "ok")
	}

	http.HandleFunc("/", respWithHeader)
	http.HandleFunc("/version", version)
	http.HandleFunc("/healthz", healthz)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
