package controllers

import (
	"net/http"
	// "database/sql"
	// "log"
	// "time"
	// "gopkg.in/gorp.v1"
	// _ "github.com/mattn/go-sqlite3"
	"github.com/labstack/echo"
)

func GetTable() echo.HandlerFunc {
	return func(c echo.Context) error {
		// dbmap := initDb()
		// defer dbmap.Db.Close()

		// err := dbmap.TruncateTables()
		// checkErr(err, "TruncateTables failed")

		return c.String(http.StatusOK, "get table!")
	}
}