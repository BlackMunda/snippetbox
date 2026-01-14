package main

import (
    "os"
    "log/slog"
    "flag"
    "net/http"
)

func main() {

    addr := flag.String("addr", ":8080", "sever starting port")
    flag.Parse()

    mux := http.NewServeMux()
    
    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))  //handleroption struct customize

    fileServer := http.FileServer(http.Dir("./ui/static/"))

    mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

    mux.HandleFunc("GET /{$}", home)
    mux.HandleFunc("GET /snippet/view/{id}", snippetView)
    mux.HandleFunc("GET /snippet/create", snippetCreate)
    mux.HandleFunc("POST /snippet/create", snippetCreatePost)

    logger.Info("starting server on ", slog.Any("addr", *addr))

    err := http.ListenAndServe(*addr, mux)
    logger.Error(err.Error())
    os.Exit(1)
}

