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
		
		dbmap := initDb()
		defer dbmap.Db.Close()
		
		var management  []HouseHold
		//SQLのクエリを直接記入（今回はselect文）
		_, err := dbmap.Select(&management, "select * from main_list")
		checkErr(err, "select failed")
		fmt.Printf("%#v", management)
		
		return c.JSON(http.StatusOK, management)
	}
}

func CreateItem() echo.HandlerFunc {
	return func(c echo.Context) error {
		
		dbmap := initDb()
		defer dbmap.Db.Close()

		var beforeItem HouseHold
		err := dbmap.SelectOne(&beforeItem,
			"select * from main_list order by Id desc limit 1")
		checkErr(err, "select failed")

		fmt.Printf("%#v\n", beforeItem)

		dateNow := time.Now()
		item := &HouseHold{beforeItem.Id + 1, dateNow, dateNow, 0, "初期データ", 4000, 1}
		err = dbmap.Insert(item)
		checkErr(err, "insert failed")

		return c.String(http.StatusOK, "created item!")
	}
}

//dbに関する設定、今後modelsに移行予定
func initDb() *gorp.DbMap {
	err := godotenv.Load()
	checkErr(err, "env failed")
	db, err := sql.Open("mysql", os.Getenv("TABLENAME"))
	checkErr(err, "sql.Open failed")
	
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF-8"}}

	// err = dbmap.CreateTablesIfNotExists()
	// checkErr(err, "Create tables failed")
	
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