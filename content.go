// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package website exports the static content as an embed.FS.
package main

import (
	"embed"
	"io/fs"
)

// TourOnly returns the content needed only for the standalone tour.
func TourOnly() fs.FS {
	return subdir(tourOnly, "content")
}

//go:embed content
var tourOnly embed.FS

func subdir(fsys fs.FS, path string) fs.FS {
	s, err := fs.Sub(fsys, path)
	if err != nil {
		panic(err)
	}
	return s
}
