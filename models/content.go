package models

import (
	"github.com/astaxie/beego/orm"
)

func GetContentOne(id int64) (*Content, error) {
	content := new(Content)
	o := orm.NewOrm()
	err := o.QueryTable("content").Filter("id", id).One(content)
	return content, err
}

func GetContentByArticleId(articleId int64) (*Content, error) {
	content := new(Content)
	o := orm.NewOrm()
	err := o.QueryTable("content").Filter("article_id", articleId).One(content)
	return content, err
}
