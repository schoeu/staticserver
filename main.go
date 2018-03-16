package main

import (
    "fmt"
	"log"
	"flag"
    "net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello golang static server!")
}

func main() {
	var prefix, path, port string
	flag.StringVar(&path, "path", "", "Path to files direction.")
	flag.StringVar(&prefix, "prefix", "", "Prefix of url path.")
	flag.StringVar(&port, "port", "8910", "Static server port.")
	flag.Parse()

	if path == "" {
		log.Fatal("Need path of files direction.")
	}


    http.HandleFunc("/", index)

    // 设置静态目录
    fsh := http.FileServer(http.Dir(path))
    http.Handle(prefix, http.StripPrefix(prefix, fsh))

    err := http.ListenAndServe(":" + port, nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}