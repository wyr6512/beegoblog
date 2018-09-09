package models

import (
	"github.com/astaxie/beego/orm"
)

func GetTagsByArticleId(articleId int64) (tags []*Tag, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("article_tag").Filter("article_id", articleId).All(&tags)
	return tags, err
}

func DeleteTags(articleId int64) error {
	o := orm.NewOrm()
	_, err := o.Raw("delete from article_tag where article_id = ?", articleId).Exec()
	return err
}
