package main

import (
    "fmt"
    "net/http"
    "flag"
    "os"
)

func main() {

    doc_root := flag.String("doc-root", "", "Document Root Directory")
    port := flag.Int("port", 0, "Port Number")
    flag.Parse()

    if (*doc_root == "" || *port == 0) {
        flag.PrintDefaults()
        return
    }

    if _, err := os.Stat(*doc_root); os.IsNotExist(err) {
        fmt.Fprintf(os.Stderr, "Document Root is not found \"%s\"\n", *doc_root)
        os.Exit(1)
    }

    fmt.Printf("Run HTTP Server for \"%s\" on localhost:%d\n", *doc_root, *port)

    http.Handle("/", http.FileServer(http.Dir(*doc_root)))
    err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %s\n", err)
        os.Exit(1)
    }
}

