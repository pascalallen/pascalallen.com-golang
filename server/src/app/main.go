package main

import (
    "fmt"
    "log"
    "database/sql"
    "net/http"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    fmt.Println("Go MySQL Tutorial")

    http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println("Hello Mars!")
    }))

    http.ListenAndServe(":3000", nil)
    log.Println("Server is now running on port 3000")

    type Post struct {
        Id string `json:"id"`
        Title string `json:"title"`
        Body string `json:"body"`
    }

    db, err := sql.Open("mysql", "homestead:secret@tcp(127.0.0.1:3306)/pascalallen.com")

    if err != nil {
        panic(err.Error())
    }

    defer db.Close()

    results, err := db.Query("SELECT * FROM posts")
    if err != nil {
        panic(err.Error())
    }

    for results.Next() {
        var post Post
        err = results.Scan(&post.Id, &post.Title, &post.Body)
        if err != nil {
            panic(err.Error())
        }
        log.Printf(post.Title)
    }

{{/*    insert, err := db.Query("INSERT INTO posts VALUES ('title', 'body')")*/}}

{{/*    if err != nil {*/}}
{{/*        panic(err.Error())*/}}
{{/*    }*/}}

{{/*    defer insert.Close()*/}}

    var post Post
    err = db.QueryRow("SELECT * FROM posts where id = ?", 1).Scan(&post.Id, &post.Title, &post.Body)
    if err != nil {
        panic(err.Error())
    }

    log.Println(post.Id)
    log.Println(post.Title)
    log.Println(post.Body)
}
