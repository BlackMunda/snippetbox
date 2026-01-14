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

    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))  //handleroption struct customize

    app := &application{
        logger : logger,
    }


    logger.Info("starting server on ", slog.Any("addr", *addr))

    err := http.ListenAndServe(*addr, app.routes())
    logger.Error(err.Error())
    os.Exit(1)
}

