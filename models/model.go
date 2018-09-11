package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"time"
)

func init() {
	orm.DefaultTimeLoc = time.UTC
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//设置默认数据库
	dbconnstr := beego.AppConfig.String("dbuser") + ":" + beego.AppConfig.String("dbpassword") + "@tcp(" + beego.AppConfig.String("dbhost") + ":" + beego.AppConfig.String("dbport") + ")/" + beego.AppConfig.String("dbname") + "?charset=utf8&loc=Local"

	orm.RegisterDataBase("default", "mysql", dbconnstr, 30)
	//注册定义的model
	orm.RegisterModel(new(Category), new(Article), new(Comment), new(Content), new(Tag), new(ArticleTag))

	// 创建table
	orm.RunSyncdb("default", false, true)
}

const (
	PageSize = 10
)

//分类
type Category struct {
	Id           int64
	Name         string    `orm:"size(32)"`
	Views        int64     `orm:"default(0)"`
	ArticleCount int64     `orm:"default(0)"`
	CreatedAt    time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt    time.Time `orm:"auto_now;type(datetime)"`
}

//文章
type Article struct {
	Id            int64
	Uid           int64
	Author        string `orm:"size(32)"`
	Title         string `orm:"size(128)"`
	Abstract      string
	CategoryId    int64
	Views         int64 `orm:"default(0)"`
	Tags          string
	ReplyCount    int64     `orm:"default(0)"`
	ReplyLastTime time.Time `orm:"auto_now;type(datetime)"`
	ReplyLastUid  int64     `orm:"default(0)"`
	Status        int       `orm:"default(0)"`
	CreatedAt     time.Time `orm:"auto_now_add;type(datetime);index"`
	UpdatedAt     time.Time `orm:"auto_now;type(datetime);index"`
}

//文章内容
type Content struct {
	Id        int64
	ArticleId int64     `orm:"index"`
	Content   string    `orm:"size(5000)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime);index"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime);index"`
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
	Name      string    `orm:"size(16);unique"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime);index"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime);index"`
}

//文章-标签对应关系
type ArticleTag struct {
	Id        int64
	ArticleId int64 `orm:"index"`
	TagId     int64 `orm:"index"`
}

//添加
func Add(obj interface{}) error {
	o := orm.NewOrm()
	_, err := o.Insert(obj)
	return err
}

//修改
func Update(obj interface{}) error {
	o := orm.NewOrm()
	_, err := o.Update(obj)
	return err
}

//删除
func Delete(obj interface{}) error {
	o := orm.NewOrm()
	_, err := o.Delete(obj)
	return err
}
