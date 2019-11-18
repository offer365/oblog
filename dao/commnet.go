package dao

import (
	"errors"

	"github.com/offer365/oblog/model"
)

func InsertComment(comment *model.Comment) (err error) {

	if comment == nil {
		err = errors.New("invalid parameter")
		return
	}

	tx, err := DB.Beginx()
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
	}()

	query := `
			INSERT INTO comment (content, username, article_id)
			VALUES (?, ?, ?)`

	_, err = tx.Exec(query, comment.Content, comment.Username, comment.ArticleId)
	if err != nil {
		return
	}

	query = `
			UPDATE article
			SET comment_count = comment_count + 1
			WHERE id = ?`

	_, err = tx.Exec(query, comment.ArticleId)
	if err != nil {
		return
	}

	err = tx.Commit()
	return
}

func UpdateViewCount(articleId int64) (err error) {

	query := `
		UPDATE article
		SET view_count = view_count + 1
		WHERE id = ?`

	_, err = DB.Exec(query, articleId)
	if err != nil {
		return
	}

	return
}

func GetCommentList(articleId int64, skip, limit int) (commentList []*model.Comment, err error) {

	if skip < 0 || limit < 0 {
		err = errors.New("invalid parameter")
		return
	}

	query := `
			SELECT id, content, username, create_at
			FROM comment
			WHERE article_id = ?
				AND status = 1
			ORDER BY create_at DESC
			LIMIT ?, ?`

	err = DB.Select(&commentList, query, articleId, skip, limit)
	return
}
