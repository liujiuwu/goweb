package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type User struct {
	Id   int64
	Name string
}

func main() {
	orm, err := xorm.NewEngine("mysql", "root:123456@/test?charset=utf8")
	defer orm.Close()
	if err != nil {
	}

	err = orm.Sync(new(User))

	orm.Insert(&User{Id: 1, Name: "John"})

}
