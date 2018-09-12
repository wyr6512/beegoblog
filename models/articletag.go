package models

import (
	"github.com/astaxie/beego/orm"
)

//根据文章id获取文章标签关联关系
func GetArticleTagsByArticleId(articleId int64) (tags []*ArticleTag, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("article_tag").Filter("article_id", articleId).All(&tags)
	return tags, err
}

//删除关联文章标签关联关系
func DeleteArticleTags(articleId int64) error {
	o := orm.NewOrm()
	_, err := o.Raw("delete from article_tag where article_id = ?", articleId).Exec()
	return err
}

//分页获取关联列表
func GetArticleTagList(id int64, pageNo, pageSize int64) (tags []*ArticleTag, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("article_tag").Filter("tag_id", id)
	offset := (pageNo - 1) * pageSize
	_, err = qs.OrderBy("-id").Limit(pageSize, offset).All(&tags)
	if err != nil {
		return nil, err
	}
	return tags, nil
}
