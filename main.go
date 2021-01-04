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

type staticHandler struct {

}

func (t staticHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token")
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")

	http.FileServer(http.Dir(dir)).ServeHTTP(w, r)
}

func main() {
	//http.Handle("/", http.FileServer(http.Dir(dir)))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), staticHandler{}); err != nil {
		panic(err)
	}
}

