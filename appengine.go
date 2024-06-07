// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

const runUrl = "https://go.dev/_/compile"

func RegisterHandlers(mux *http.ServeMux) error {
	prepContent = gaePrepContent
	socketAddr = gaeSocketAddr
	analyticsHTML = template.HTML(os.Getenv("TOUR_ANALYTICS"))

	if err := initTour(mux, "HTTPTransport"); err != nil {
		return err
	}

	return nil
}

// gaePrepContent returns a Reader that produces the content from the given
// Reader, but strips the prefix "#appengine:", optionally followed by a space, from each line.
// It also drops any non-blank line that follows a series of 1 or more lines with the prefix.
func gaePrepContent(in io.Reader) io.Reader {
	var prefix = []byte("#appengine:")
	out, w := io.Pipe()
	go func() {
		r := bufio.NewReader(in)
		drop := false
		for {
			b, err := r.ReadBytes('\n')
			if err != nil && err != io.EOF {
				w.CloseWithError(err)
				return
			}
			if bytes.HasPrefix(b, prefix) {
				b = b[len(prefix):]
				if b[0] == ' ' {
					// Consume a single space after the prefix.
					b = b[1:]
				}
				drop = true
			} else if drop {
				if len(b) > 1 {
					b = nil
				}
				drop = false
			}
			if len(b) > 0 {
				w.Write(b)
			}
			if err == io.EOF {
				w.Close()
				return
			}
		}
	}()
	return out
}

func runHandler(w http.ResponseWriter, r *http.Request) {
	if err := passThru(w, r); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Compile server error.")
	}
}

func passThru(w io.Writer, req *http.Request) error {
	defer req.Body.Close()
	req.Header.Set("User-Agent", "go-vim")
	r, err := http.Post(runUrl, req.Header.Get("Content-type"), req.Body)
	if err != nil {
		log.Fatalf("making POST request: %v", err)
		return err
	}
	defer r.Body.Close()
	if _, err := io.Copy(w, r.Body); err != nil {
		log.Fatalf("copying response Body: %v", err)
		return err
	}
	return nil
}

// gaeSocketAddr returns the WebSocket handler address.
// The App Engine version does not provide a WebSocket handler.
func gaeSocketAddr() string { return "" }
