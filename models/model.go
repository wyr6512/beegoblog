package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"time"
)

func init() {
	//注册驱动
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	//设置默认数据库
	orm.RegisterDataBase("default", "mysql", "root:123456@/blog?charset=utf8", 30)
	//注册定义的model
	orm.RegisterModel(new(Category), new(Article), new(Comment))

	// 创建table
	orm.RunSyncdb("default", false, true)
}

type Category struct {
}

type Article struct {
}

type Comment struct {
}

type Tag struct {
}
