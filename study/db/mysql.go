package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type MySQL struct {
	db  *sql.DB
	err error
}

func (mysql *MySQL) Conn() {
	var config Config
	mysql.db, mysql.err = sql.Open("mysql", config.dsn())
	mysql.err = mysql.db.Ping()
	check(mysql.err)
}

func (mysql *MySQL) GetAll(sql string, params ...interface{}) []map[string]string {
	rows, err := mysql.db.Query(sql, params...)
	check(err)

	columns, _ := rows.Columns()
	//这里表示一行所有列的值，用[]byte表示
	values := make([][]byte, len(columns))
	//这里表示一行填充数据
	scans := make([]interface{}, len(columns))
	//这里scans引用values，把数据填充到[]byte里
	for k := range values {
		scans[k] = &values[k]
	}

	i := 0
	result := make([]map[string]string, 0)
	for rows.Next() {
		rows.Scan(scans...)

		row := make(map[string]string)
		for k, v := range values {
			key := columns[k]
			//这里把[]byte数据转成string
			row[key] = string(v)
		}

		result = append(result, row)
		i++
	}

	return result
}

func (mysql *MySQL) Exec(sql string, params ...interface{}) (row int64, err error) {
	row, err = mysql.Exec(sql, params...)
	if err != nil {
		log.Printf("mysql.Exec sql(%s) params(%+v) error(%+v)\n", sql, params, err)
	}
	return
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
