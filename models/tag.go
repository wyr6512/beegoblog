package models

import (
	"github.com/astaxie/beego/orm"
)

func GetTagOne(id int64) (*Tag, error) {
	tag := new(Tag)
	o := orm.NewOrm()
	err := o.QueryTable("tag").Filter("id", id).One(tag)
	return tag, err
}

//获取列表
func GetTagList(name string, pageNo, pageSize int64) (tags []*Tag, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("tag")
	if len(name) > 0 {
		qs = qs.Filter("name__contains", name)
	}
	offset := (pageNo - 1) * pageSize
	_, err = qs.OrderBy("-id").Limit(pageSize, offset).All(&tags)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func GetTagsByNames(tagNames []string) (tags []*Tag, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("tag")
	_, err = qs.Filter("name__in", tagNames).All(&tags)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

//获取总数
func GetTagCount(name string) (int64, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("tag")
	if len(name) > 0 {
		qs = qs.Filter("name__contains", name)
	}
	cnt, err := qs.Count()
	if err != nil {
		return 0, err
	}
	return cnt, nil
}
