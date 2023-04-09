package main

import (
	"fmt"

	"my-go/study/db"
)

func main() {
	d := db.MySQL{}
	d.Conn()
	result := d.GetAll("select * from py_house_price where id >= ? order by id desc limit ?", 105, 10)

	for k, v := range result {
		fmt.Printf("key: %d, id: %s, name: %s, last_price: %s, current_price: %s\n",
			k,
			v["id"],
			v["name"],
			v["last_week_price"],
			v["current_week_price"],
		)
	}
}
