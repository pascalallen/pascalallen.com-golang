package mysql

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

var Client, err := sql.Open("mysql", "homestead:secret@127.0.0.1:3306/pascalallen.com")
