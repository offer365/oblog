package model

import "time"

// 文章
type ArticleInfo struct {
	Id           int64     `db:"id"`
	CategoryId   int64     `db:"category_id"`
	Summary      string    `db:"summary"`
	Title        string    `db:"title"`
	ViewCount    uint32    `db:"view_count"`
	CreateAt     time.Time `db:"create_at"`
	CommentCount uint32    `db:"comment_count"`
	Username     string    `db:"username"`
}

// 文章详情 包含分类信息
type ArticleDetail struct {
	ArticleInfo
	Content string `db:"content"`
	Category
}

// 文章列表 包含分类信息和文章信息
type ArticleRecord struct {
	ArticleInfo
	Category
}

// 文章分类
type Category struct {
	Id     int64  `db:"id"`
	Name   string `db:"name"`
	Number int    `db:"number"`
}

// 评论
type Comment struct {
	Id        int64     `db:"id"`
	Content   string    `db:"content"` // 评论内容
	Username  string    `db:"username"`
	CreateAt  time.Time `db:"create_at"`
	Status    int       `db:"status"`
	ArticleId int64     `db:"article_id"` // 评论文章
}

// 留言
type Leave struct {
	Id       int64     `db:"id"`
	Content  string    `db:"content"`
	Username string    `db:"username"`
	CreateAt time.Time `db:"create_at"`
	Email    string    `db:"email"`
}

// 相关文章
type RelativeArticle struct {
	Id    int64  `db:"id"`
	Title string `db:"title"`
}
