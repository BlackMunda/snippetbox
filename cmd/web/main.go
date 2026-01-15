package main

import (
    _ "github.com/go-sql-driver/mysql"
    "snippetbox.net/internal/models"
    "database/sql"
    "os"
    "log/slog"
    "flag"
    "net/http"
)

type application struct {
    logger *slog.Logger
    snippets *models.SnippetModel
}

func main() {

    addr := flag.String("addr", ":8080", "sever starting port")
    dsn := flag.String("dsn", "web:Dev@2005@/snippetbox?parseTime=true", "MySQL data source name")
    flag.Parse()

    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))  //handleroption struct customize

    db, err := openDb(*dsn)
    if err != nil {
        logger.Error(err.Error())
        os.Exit(1)
    }
    defer db.Close()

    app := &application{
        logger : logger,
        snippets : &models.SnippetModel{
            DB : db,
        },
    }

    logger.Info("starting server on ", slog.Any("addr", *addr))

    err = http.ListenAndServe(*addr, app.routes())
    logger.Error(err.Error())
    os.Exit(1)
}

func openDb(dsn string) (*sql.DB, error) {
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }

    err = db.Ping()
    if err != nil {
        db.Close()
        return nil, err
    }

    return db, nil
}

