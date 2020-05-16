package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	port int
	dir string
)

func init () {
	flag.IntVar(&port, "p", 80, "port")
	flag.StringVar(&dir, "d", ".", "directory")
	flag.Parse()
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(dir)))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}

