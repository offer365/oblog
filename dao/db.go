package dao

import (
	"github.com/jmoiron/sqlx"
	"github.com/offer365/example/mysql"
)


var (
	DB *sqlx.DB
)

func Init(addr, username, password, database, char string, parse bool) (err error) {
	db := mysql.NewDB("mysql")
	DB, err = db.Init(
		mysql.WithAddr(addr),
		mysql.WithUsername(username),
		mysql.WithDatabase(database),
		mysql.WithPassword(password),
		mysql.WithCharSet(char),
		mysql.WithMaxConn(100),
		mysql.WithMaxIdleConn(10),
		mysql.WithParseTime(parse),
	)
	return nil
}
