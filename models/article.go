package models

import (
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strings"
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

func AddArticle(article *Article, content string) error {
	orm.Debug = true
	o := orm.NewOrm()
	o.Begin()
	tags := article.Tags
	var tagsArr []string
	if len(tags) > 0 { //标签处理
		tagsArr = strings.Split(tags, ",")
		tags = strings.Join(tagsArr, "$#")
		tags = "#" + tags + "$"
	}
	article.Tags = tags
	_, err := o.Insert(article)
	if err != nil {
		err = o.Rollback()
	}
	//添加文章内容
	cont := &Content{
		Content:   content,
		ArticleId: article.Id,
	}
	_, err = o.Insert(cont)
	if err != nil {
		err = o.Rollback()
	}
	existsTags, err := GetTagsByNames(tagsArr)
	if err != nil {
		err = o.Rollback()
	}
	mapTags := ArrayToMap(existsTags)
	for _, tag := range tagsArr {
		var tagModel *Tag
		if m, ok := mapTags[tag]; ok {
			tagModel = m
		} else { //标签不存在，添加
			t := &Tag{
				Name: tag,
			}
			_, err := o.Insert(t)
			if err != nil {
				err = o.Rollback()
			}
			tagModel = t
		}
		//添加标签和文章对应关系
		articleTag := &ArticleTag{
			ArticleId: article.Id,
			TagId:     tagModel.Id,
		}
		_, err = o.Insert(articleTag)
		if err != nil {
			err = o.Rollback()
		}
	}
	o.Commit()
	return err
}

func UpdateArticle(article *Article, content *Content) error {
	o := orm.NewOrm()
	o.Begin()
	tags := article.Tags
	var tagsArr []string
	if len(tags) > 0 { //标签处理
		tagsArr = strings.Split(tags, ",")
		tags = strings.Join(tagsArr, "$#")
		tags = "#" + tags + "$"
	}
	article.Tags = tags
	_, err := o.Update(article)
	if err != nil {
		err = o.Rollback()
	}
	//更新文章内容
	_, err = o.Update(content)
	if err != nil {
		err = o.Rollback()
	}
	existsTags, err := GetTagsByNames(tagsArr)
	if err != nil {
		err = o.Rollback()
	}
	mapTags := ArrayToMap(existsTags)
	//删除已存在文章标签关系
	err = DeleteTags(article.Id)
	if err != nil {
		err = o.Rollback()
	}
	for _, tag := range tagsArr {
		var tagModel *Tag
		if m, ok := mapTags[tag]; ok {
			tagModel = m
		} else { //标签不存在，添加
			t := &Tag{
				Name: tag,
			}
			_, err := o.Insert(t)
			if err != nil {
				err = o.Rollback()
			}
			tagModel = t
		}
		//添加标签和文章对应关系
		articleTag := &ArticleTag{
			ArticleId: article.Id,
			TagId:     tagModel.Id,
		}
		_, err = o.Insert(articleTag)
		if err != nil {
			err = o.Rollback()
		}
	}
	o.Commit()
	return err
}

func DeleteArticle(article *Article) error {
	o := orm.NewOrm()
	o.Begin()
	//删除文章内容
	_, err := o.Raw("delete from content where article_id = ?", article.Id).Exec()
	if err != nil {
		err = o.Rollback()
	}
	_, err = o.Delete(article)
	if err != nil {
		err = o.Rollback()
	}
	o.Commit()
	return err
}

//数组转map
func ArrayToMap(tags []*Tag) map[string]*Tag {
	var mTags = make(map[string]*Tag)
	for _, t := range tags {
		mTags[t.Name] = t
	}
	return mTags
}
