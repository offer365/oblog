package dao

import (
	"errors"
	"fmt"

	"../model"
)

func InsertArticle(article *model.ArticleDetail) (articleId int64, err error) {
	if article == nil {
		err = errors.New("invalid article parameter")
		return
	}

	query := `
			INSERT INTO article (content, summary, title, username, category_id
				, view_count, comment_count)
			VALUES (?, ?, ?, ?, ?
				, ?, ?)`

	result, err := DB.Exec(query, article.Content, article.Summary,
		article.Title, article.Username, article.ArticleInfo.CategoryId,
		article.ArticleInfo.ViewCount, article.ArticleInfo.CommentCount)
	if err != nil {
		return
	}

	articleId, err = result.LastInsertId()
	return
}

func GetArticleList(skip, limit int) (articleList []*model.ArticleInfo, err error) {

	if skip < 0 || limit < 0 {
		err = errors.New("invalid article parameter")
		return
	}

	query := `
			SELECT id, summary, title, view_count, create_at
				, comment_count, username, category_id
			FROM article
			WHERE status = 1
			ORDER BY create_at DESC
			LIMIT ?, ?`

	err = DB.Select(&articleList, query, skip, limit)
	return
}

func GetArticleListByCategoryId(categoryId, skip, limit int) (articleList []*model.ArticleInfo, err error) {

	if skip < 0 || limit < 0 {
		err = errors.New("invalid article parameter")
		return
	}

	query := `
			SELECT id, summary, title, view_count, create_at
				, comment_count, username, category_id
			FROM article
			WHERE status = 1
				AND category_id = ?
			ORDER BY create_at DESC
			LIMIT ?, ?`

	err = DB.Select(&articleList, query, categoryId, skip, limit)
	return
}

func GetArticleDetail(articleId int64) (articleDetail *model.ArticleDetail, err error) {

	if articleId < 0 {
		err = fmt.Errorf("invalid parameter,article_id:%d", articleId)
		return
	}

	articleDetail = &model.ArticleDetail{}
	query := `
			SELECT id, summary, title, view_count, content
				, create_at, comment_count, username, category_id
			FROM article
			WHERE id = ?
				AND status = 1`

	err = DB.Get(articleDetail, query, articleId)
	return
}

// 获取相关文章 获取与指定文章相同分类的其他文章。
func GetRelativeArticle(articleId int64) (articleList []*model.RelativeArticle, err error) {
	query := `
			SELECT id, title
			FROM article
			WHERE category_id = (
					SELECT category_id
					FROM article
					WHERE id = ?
				)
				AND id != ?
			LIMIT 10`
	err = DB.Select(&articleList, query, articleId, articleId)
	return
}

// 前一个
func GetPrevArticleById(articleId int64) (info *model.RelativeArticle, err error) {

	info = &model.RelativeArticle{
		Id: -1,
	}
	query := `
			SELECT id, title
			FROM article
			WHERE id < ?
			ORDER BY id DESC
			LIMIT 1`
	err = DB.Get(info, query, articleId)
	if err != nil {
		return
	}

	return
}

// 后一个
func GetNextArticleById(articleId int64) (info *model.RelativeArticle, err error) {
	info = &model.RelativeArticle{
		Id: -1,
	}
	query := `
			SELECT id, title
			FROM article
			WHERE id > ?
			ORDER BY id ASC
			LIMIT 1`
	err = DB.Get(info, query, articleId)
	if err != nil {
		return
	}

	return
}

func IsArticleExist(articleId int64) (exists bool, err error) {
	var n = 0
	query := `
			SELECT COUNT(id) AS n
			FROM article
			WHERE id = ?`
	err = DB.Get(&n, query, articleId)
	if n > 0 {
		exists = true
	}
	return
}
