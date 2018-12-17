package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	var filePath, port, help string
	flag.StringVar(&filePath, "path", "", "Path to files direction.")
	flag.StringVar(&port, "port", "8910", "Static server port.")
	flag.StringVar(&help, "help", "", "help info for staticserver")
	flag.Parse()

	if filePath == "" {
		fmt.Println("Need path of files direction.")
		return
	}

	fmt.Printf("Server on http://127.0.0.1:%s\n", port)

	fsh := http.FileServer(http.Dir(filePath))
	http.Handle("/", fsh)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
