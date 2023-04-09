package components

import (
	"my-go/study/db"
)

var Conn db.MySQL

func Init() {
	Conn = db.MySQL{}
	Conn.Conn()
}
