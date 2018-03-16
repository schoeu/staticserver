package main

import (
    "fmt"
	"log"
	"flag"
    "net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I'm ready for that, you know.")
}

func main() {
	var prefix, filePath, port, help string
	flag.StringVar(&filePath, "path", "", "Path to files direction.")
	flag.StringVar(&prefix, "prefix", "/static/", "Prefix of url path.")
	flag.StringVar(&port, "port", "8910", "Static server port.")
	flag.StringVar(&help, "help", "", "help info for staticserver")
	flag.Parse()

	if filePath == "" {
		log.Fatal("Need path of files direction.")
	}

	// index page.
    http.HandleFunc("/", index)

    fsh := http.FileServer(http.Dir(filePath))
    http.Handle(prefix, http.StripPrefix(prefix, fsh))

    err := http.ListenAndServe(":" + port, nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
