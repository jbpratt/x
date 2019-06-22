package main

import (
	"log"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("./html"))
	err := http.ListenAndServe(":8000",
		http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			if req.URL.Path == "/test.wasm" {
				resp.Header().Set("content-type", "application/wasm")
			}

			fs.ServeHTTP(resp, req)
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
}
