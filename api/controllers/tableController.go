package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"gopkg.in/gorp.v2"
	_ "github.com/mattn/go-sqlite3"
	"github.com/labstack/echo"
)

type HouseHold struct {
	Id      int64  `db:"post_id"`
	Created int64
	Title   string `db:",size:50"`               
	Body    string `db:"article_body,size:1024"` 
}

func GetTable() echo.HandlerFunc {
	return func(c echo.Context) error {
		dbmap := initDb()
		defer dbmap.Db.Close()

		err := dbmap.TruncateTables()
		checkErr(err, "TruncateTables failed")

		var management  HouseHold
		//SQLのクエリを直接記入（今回はselect文）
		err = dbmap.SelectOne(&management, "select * from household where id=?", 1)

		return c.String(http.StatusOK, "get table!")
	}
}

//dbに関する設定、今後modelsに移行予定
func initDb() *gorp.DbMap {
	db, err := sql.Open("mysql", "/example")
	checkErr(err, "sql.Open failed")

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF-8"}}

	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")
	
	dbmap.AddTableWithName(HouseHold{}, "household").SetKeys(true, "Id")
	fmt.Printf("%v", management)

	return dbmap
}

//エラーハンドリング用関数
func checkErr(err error, msg string) {
	if err != nil {
			log.Fatalln(msg, err)
	}
}