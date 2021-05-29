package models
import "time"


type ArticleData struct {
	Title string `binding:"required"`
	Description string `binding:"required"`
	Body string `binding:"required"`
        TagList []string
}

type ArticleMeta struct {
	Slug string `binding:"required"`
        CreatedAt *time.Time
}

type ArticleDetail struct {
        ArticleData
        ArticleMeta
}
