package res

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sluu99/uuid"
	"strings"
)

var db orm.Ormer

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:admin@tcp(localhost:3306)/gobus")
	orm.Debug = true

}

func Orm_demo_(contains interface{}) {
	db := orm.NewOrm()
	db.Using("default")
	//db.Raw("select name,id from demo_test where name like ?", "%sanli%").QueryRows(contains)
	var count int
	db.Raw("select count(1) from topic").QueryRow(&count)
	fmt.Println(count)
}

func Orm_demo() {
	var users []user
	Orm_demo_(&users)
	for _, value := range users {
		fmt.Println(value)
	}
}

type user struct {
	Id   string
	Name string
}

func InsertSingleTopic(topic Topic) {
	db := orm.NewOrm()
	db.Using("default")
	id := strings.ToUpper(strings.Replace(uuid.Rand().Hex(), "-", "", -1))
	db.Raw(sql_0001, id, topic.Title, topic.Content).Exec()
}

type Topic struct {
	Id      string
	Title   string
	Content string
}
