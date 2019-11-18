package dao

import (
	"../model"
	"github.com/jmoiron/sqlx"
)

func InsertCategory(category *model.Category) (categoryId int64, err error) {

	query := `insert into category(name, number)values(?,?)`
	result, err := DB.Exec(query, category.Name, category.Number)
	if err != nil {
		return
	}

	categoryId, err = result.LastInsertId()
	return
}

func GetCategoryList(categoryIds []int64) (categoryList []*model.Category, err error) {
	// In扩展args中的slice值，返回修改后的查询字符串 和可以由数据库执行的新arg列表。
	query := `select id, name, number from category where id in(?)`
	query, args, err := sqlx.In(query, categoryIds)
	if err != nil {
		return
	}

	err = DB.Select(&categoryList, query, args...)
	return
}

func GetAllCategoryList() (categoryList []*model.Category, err error) {

	query := `select id, name, number from category order by number asc`
	err = DB.Select(&categoryList, query)
	return
}

func GetCategoryById(id int64) (category *model.Category, err error) {

	category = &model.Category{}
	query := `select id, name, number from category where id=?`
	err = DB.Get(category, query, id)
	return
}
