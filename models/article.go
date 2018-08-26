package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type ArticleParams struct {
	Title    string
	Status   int
	FromTime string
	EndTime  string
	Tag      string
}

func GetArticleOne(id int64) (*Article, error) {
	article := new(Article)
	o := orm.NewOrm()
	err := o.QueryTable("article").One(article)
	return article, err
}

func GetArticleList(params ArticleParams, pageNo, pageSize int64) (article []*Article, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("article")
	if len(params.Title) > 0 {
		qs = qs.Filter("title__contains", params.Title)
	}
	if params.Status > 0 {
		qs = qs.Filter("status", params.Status)
	}
	if len(params.FromTime) > 0 {
		t, err := time.Parse("2006-01-02 15:04:05", params.FromTime)
		if err == nil {
			qs = qs.Filter("created_at__gt", t)
		}
	}
	if len(params.EndTime) > 0 {
		t, err := time.Parse("2006-01-02 15:04:05", params.EndTime)
		if err == nil {
			qs = qs.Filter("created_at__lte", t)
		}
	}
	if len(params.Tag) > 0 {
		qs = qs.Filter("tags__contains", params.Tag)
	}
	offset := (pageNo - 1) * pageSize
	_, err = qs.OrderBy("-id").Limit(pageSize, offset).All(&article)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func GetArticleCount(params ArticleParams) (int64, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("article")
	if len(params.Title) > 0 {
		qs = qs.Filter("title__contains", params.Title)
	}
	if params.Status > 0 {
		qs = qs.Filter("status", params.Status)
	}
	if len(params.FromTime) > 0 {
		t, err := time.Parse("2006-01-02 15:04:05", params.FromTime)
		if err == nil {
			qs = qs.Filter("created_at__gt", t)
		}
	}
	if len(params.EndTime) > 0 {
		t, err := time.Parse("2006-01-02 15:04:05", params.EndTime)
		if err == nil {
			qs = qs.Filter("created_at__lte", t)
		}
	}
	if len(params.Tag) > 0 {
		qs = qs.Filter("tags__contains", params.Tag)
	}
	cnt, err := qs.Count()
	if err != nil {
		return 0, err
	}
	return cnt, nil
}
