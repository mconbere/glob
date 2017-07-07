package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func init() {
	files, err := filepath.Glob("static/*.html")
	if err != nil {
		panic(fmt.Errorf("filepath.Glob failed: %v", err))
	}
	for _, f := range files {
		http.HandleFunc("/"+f, func(w http.ResponseWriter, r *http.Request) {
			fr, err := os.Open(f)
			if err != nil {
				panic(err)
			}
			io.Copy(w, fr)
		})
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Got files: %v", files)
	})
}
