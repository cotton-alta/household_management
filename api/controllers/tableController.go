package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"gopkg.in/gorp.v2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/joho/godotenv"
)

type HouseHold struct {
	Id      int64  `db:"id, primarykey"`
	Created time.Time  `db:"created_at"`
	Updated time.Time  `db:"updated_at"`
	Amount  int64  `db:"amount_money"`
	Body    string `db:"remark,size:1024"`
	Balance int64  `db:"balance"` 
	Genre int64  `db:"genre"` 
}

func GetTable() echo.HandlerFunc {
	return func(c echo.Context) error {
		err := godotenv.Load()
		checkErr(err, "env failed")

		dbmap := initDb()
		defer dbmap.Db.Close()

		err = dbmap.TruncateTables()
		checkErr(err, "TruncateTables failed")

		dateNow := time.Now()
		log1 := &HouseHold{0, dateNow, dateNow, 0, "初期データ", 3000, 1}
		err = dbmap.Insert(log1)

		var management  []HouseHold
		//SQLのクエリを直接記入（今回はselect文）
		_, err = dbmap.Select(&management, "select * from main_list")
		// list, _ := dbmap.Select(HouseHold{}, "select * from main_list")
		fmt.Printf("%#v", management)

		return c.JSON(http.StatusOK, management)
	}
}

//dbに関する設定、今後modelsに移行予定
func initDb() *gorp.DbMap {
	db, err := sql.Open("mysql", os.Getenv("TABLENAME"))
	checkErr(err, "sql.Open failed")

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF-8"}}

	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")
	
	table := dbmap.AddTableWithName(HouseHold{}, "main_list")
	table.ColMap("Id").Rename("id")
	table.ColMap("Created").Rename("created_at")
	table.ColMap("Updated").Rename("updated_at")
	table.ColMap("Amount").Rename("amount_money")
	table.ColMap("Body").Rename("remark")
	table.ColMap("Balance").Rename("balance")
	table.ColMap("Genre").Rename("genre")

	return dbmap
}

//エラーハンドリング用関数
func checkErr(err error, msg string) {
	if err != nil {
			log.Fatalln(msg, err)
	}
}