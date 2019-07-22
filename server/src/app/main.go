package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    fmt.Println("Go MySQL Tutorial")

    db, err := sql.Open("mysql", "homestead:secret@tcp(127.0.0.1:3306)/pascalallen.com")

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

    defer db.Close()

    // perform a db.Query insert
    insert, err := db.Query("INSERT INTO posts VALUES (1, 'title', 'body')")

    // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }

    // be careful deferring Queries if you are using transactions
    defer insert.Close()
}
