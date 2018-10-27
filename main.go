package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	var prefix, filePath, port, help string
	flag.StringVar(&filePath, "path", "", "Path to files direction.")
	flag.StringVar(&prefix, "prefix", "", "Prefix of url path.")
	flag.StringVar(&port, "port", "8910", "Static server port.")
	flag.StringVar(&help, "help", "", "help info for staticserver")
	flag.Parse()

	if filePath == "" {
		fmt.Println("Need path of files direction.")
		return
	}

	fsh := http.FileServer(http.Dir(filePath))
	http.Handle("/", fsh)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
