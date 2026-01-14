package main

import (
    "os"
    "log/slog"
    "flag"
    "net/http"
)

type application struct {
    logger *slog.Logger
}

func main() {

    addr := flag.String("addr", ":8080", "sever starting port")
    flag.Parse()

    mux := http.NewServeMux()

    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))  //handleroption struct customize

    app := &application{
        logger : logger,
    }


    fileServer := http.FileServer(http.Dir("./ui/static/"))

    mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
    mux.HandleFunc("GET /{$}", app.home)
    mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
    mux.HandleFunc("GET /snippet/create", app.snippetCreate)
    mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

    logger.Info("starting server on ", slog.Any("addr", *addr))

    err := http.ListenAndServe(*addr, mux)
    logger.Error(err.Error())
    os.Exit(1)
}

