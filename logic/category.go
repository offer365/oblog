package logic

import (
	"fmt"

	"github.com/offer365/oblog/dal/db"
	"github.com/offer365/oblog/model"
)

func GetAllCategoryList() (categoryList []*model.Category, err error) {
	// 1. 从数据库中，获取文章分类列表
	categoryList, err = db.GetAllCategoryList()
	if err != nil {
		fmt.Printf("1 get article list failed, err:%v\n", err)
		return
	}

	return
}

