package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"time"
)

func init() {
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//设置默认数据库
	dbconnstr := beego.AppConfig.String("dbuser") + ":" + beego.AppConfig.String("dbpassword") + "@tcp(" + beego.AppConfig.String("dbhost") + ":" + beego.AppConfig.String("dbport") + ")/" + beego.AppConfig.String("dbname") + "?charset=utf8"

	orm.RegisterDataBase("default", "mysql", dbconnstr, 30)
	//注册定义的model
	orm.RegisterModel(new(Category), new(Article), new(Comment), new(Content), new(Tag), new(ArticleTag))

	// 创建table
	orm.RunSyncdb("default", false, true)
}

const (
	PageSize = 10
)

type Pager struct {
	PageNo     int64
	PageSize   int64
	TotalPage  int64
	TotalCount int64
	List       interface{}
}

//分类
type Category struct {
	Id           int64
	Name         string `orm:"size(32)"`
	Views        int64
	ArticleCount int64
	CreatedAt    time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateAt     time.Time `orm:"auto_now;type(datetime)"`
}

//文章
type Article struct {
	Id            int64
	Uid           int64
	Author        string `orm:"size(32)"`
	Title         string `orm:"size(128)"`
	Abstract      string
	CategoryId    int64
	Views         int64
	Tags          string
	ReplyCount    int64
	ReplyLastTime time.Time `orm:"auto_now;type(datetime)"`
	ReplyLastUid  int64
	Status        int       `orm:"default(0)"`
	CreatedAt     time.Time `orm:"auto_now_add;type(datetime);index"`
	UpdateAt      time.Time `orm:"auto_now;type(datetime);index"`
}

//文章内容
type Content struct {
	Id        int64
	ArticleId int64     `orm:"index"`
	Content   string    `orm:"size(5000)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime);index"`
	UpdateAt  time.Time `orm:"auto_now;type(datetime);index"`
}

//评论
type Comment struct {
	Id        int64
	ArticleId int64     `orm:"index"`
	Name      string    `orm:"size(32)"`
	Content   string    `orm:"size(512)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime);index"`
}

//标签
type Tag struct {
	Id        int64
	Name      string    `orm:"size(16)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime);index"`
}

//文章-标签对应关系
type ArticleTag struct {
	Id        int64
	ArticleId int64 `orm:"index"`
	TagId     int64 `orm:"index"`
}
