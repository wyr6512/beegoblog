package models

import (
	"github.com/astaxie/beego/orm"
)

//添加
func AddCategory(cate *Category) error {
	o := orm.NewOrm()
	_, err := o.Insert(cate)
	return err
}

//修改
func UpdateCategory(cate *Category) error {
	o := orm.NewOrm()
	_, err := o.Update(cate)
	return err
}

//删除
func DeleteCategory(cate *Category) error {
	o := orm.NewOrm()
	_, err := o.Delete(cate)
	return err
}

//获取单个
func GetCategoryOne(id int64) (*Category, error) {
	cate := new(Category)
	o := orm.NewOrm()
	err := o.QueryTable("category").Filter("id", id).One(cate)
	return cate, err
}

//获取列表
func GetCategoryList(name string, pageNo, pageSize int64) (cates []*Category, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("category")
	if len(name) > 0 {
		qs = qs.Filter("name__contains", name)
	}
	offset := (pageNo - 1) * pageSize
	_, err = qs.OrderBy("-id").Limit(pageSize, offset).All(&cates)
	if err != nil {
		return nil, err
	}
	return cates, nil
}

//获取总数
func GetCategoryCount(name string) (int64, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("category")
	if len(name) > 0 {
		qs = qs.Filter("name__contains", name)
	}
	cnt, err := qs.Count()
	if err != nil {
		return 0, err
	}
	return cnt, nil
}
