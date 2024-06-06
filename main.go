// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	if os.Getenv("GAE_ENV") == "standard" {
		log.Println("running in App Engine Standard mode")
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
		siteMux := http.NewServeMux()
		if err := RegisterHandlers(siteMux); err != nil {
			log.Fatal(err.Error())
		}
		http.HandleFunc("/_/fmt", fmtHandler)
		http.HandleFunc("/_/compile", runHandler)
		http.Handle("/", siteMux)
		siteMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/tour/", http.StatusFound)
		})

		log.Fatal(http.ListenAndServe(":"+port, nil))
		return
	}
	Main()
}
