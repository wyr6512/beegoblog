package models

import (
	"github.com/astaxie/beego/orm"
)

func AddCategory(cate *Category) error {
	o := orm.NewOrm()
	_, err := o.Insert(cate)
	return err
}

func UpdateCategory(cate *Category) error {
	o := orm.NewOrm()
	_, err := o.Update(cate)
	return err
}

func DeleteCategory(cate *Category) error {
	o := orm.NewOrm()
	_, err := o.Delete(cate)
	return err
}

func GetCategoryOne(cate *Category) error {
	o := orm.NewOrm()
	err := o.QueryTable("category").One(cate)
	return err
}
func GetCategoryList(name string, pageNo, pageSize int64) (Pager, error) {
	pager := Pager{
		PageNo:   pageNo,
		PageSize: pageSize,
	}
	o := orm.NewOrm()
	qs := o.QueryTable("category")
	if len(name) > 0 {
		qs.Filter("name", name)
	}
	cnt, err := qs.Count()
	if err != nil {
		return pager, err
	}
	pager.TotalCount = cnt
	pager.TotalPage = cnt / pageSize
	if cnt%pageSize > 0 {
		pager.TotalPage = cnt/pageSize + 1
	}
	offset := (pageNo - 1) * pageSize
	list := qs.Limit(pageSize, offset)
	pager.List = list
	return pager, nil
}
