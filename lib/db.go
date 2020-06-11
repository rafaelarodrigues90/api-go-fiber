package lib

import (
	"log"

	"upper.io/db.v3"
	"upper.io/db.v3/mysql"
)

var config = mysql.ConnectionURL {
	Host: "localhost",
	User: "root",
	Password: "",
	Database: "membros",
}

// Session faz conexão com o bd
var Session db.Database

func init() {
	var err error

	Session, err = mysql.Open(config)
	if err != nil {
		log.Fatal(err.Error())
	}
}