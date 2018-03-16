package main

import (
    "fmt"
	"flag"
	"net/http"
	"strings"
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
		fmt.Println("Need path of files direction.")
		return
	}

	// index page.
    http.HandleFunc("/", index)

	fsh := http.FileServer(http.Dir(filePath))
	prefix = pathProcess(prefix)
    http.Handle(prefix, http.StripPrefix(prefix, fsh))

    err := http.ListenAndServe(":" + port, nil)
    if err != nil {
        fmt.Println("ListenAndServe: ", err)
    }
}

// trans prefix
func pathProcess(p string) string{
	purePrefix := strings.Replace(p, "/", "", -1)
	return "/" + purePrefix + "/"
}
