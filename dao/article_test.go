package dao

import (
	"fmt"
	"testing"
	"time"

	"github.com/offer365/oblog/model"
)

func TestInsertArticle(t *testing.T) {
	err := Init("117.107.146.194:6603", "oblog", "qwe123!@#QWE", "oblog", "utf8mb4", true)
	if err != nil {
		t.Fatal(err)
	}
	article := &model.ArticleDetail{}
	article.ArticleInfo.CategoryId = 1
	article.ArticleInfo.CommentCount = 0
	article.Content = "this a test ak dkdkdkddkddkd"
	article.ArticleInfo.CreateAt = time.Now()
	article.ArticleInfo.Summary = `使用mysql的时间字段遇到如下两个问题
	1.使用go-sql-driver来连接mysql数据库，获取的时区默认是UTC +0的，与本地的东八区是有区别，在业务处理中会出现问题
	2.获取mysql中的日期，是string类型，需要在代码中用time.Parse进行转化`
	article.ArticleInfo.Title = "GOLANG 连接Mysql的时区问题"
	article.ArticleInfo.Username = "少林之巅"
	article.ArticleInfo.ViewCount = 1
	article.Category.Id = 1

	articleId, err := InsertArticle(article)
	if err != nil {
		t.Errorf("insert article failed, err:%v\n", err)
		return
	}

	t.Logf("insert article succ, articleId:%d\n", articleId)
	// InsertArticle(article *model.ArticleDetail) (articleId int64, err error)
}

func TestGetRelativeArticle(t *testing.T) {
	err := Init("117.107.146.194:6603", "oblog", "qwe123!@#QWE", "oblog", "utf8mb4", true)
	if err != nil {
		t.Fatal(err)
	}

	as, err := GetRelativeArticle(1)
	fmt.Println(as[0].Id, as[0].Title)
	as, err = GetRelativeArticle(2)
	fmt.Println(as[0].Id, as[0].Title)
}
