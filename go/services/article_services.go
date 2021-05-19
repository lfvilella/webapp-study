package services

import "webapp-go/models"

func ArticleCreate(data models.ArticleData) models.ArticleDetail {
	article := models.ArticleDetail{
                ArticleData: data,
                ArticleMeta: models.ArticleMeta {
                        Slug: "how-to-train-your-dragon",
                        CreatedAt: nil,
                },
	}
	return article
}

func ArticleDetail(Slug string) models.ArticleDetail {
	article := models.ArticleDetail{
                models.ArticleData {
                        Title: "How to train your dragon",
                        Description: "Ever wonder how?",
                        Body: "It takes a Jacobian",
                        TagList: []string{"dragons", "training"},
                },
                models.ArticleMeta {
                        Slug: Slug,
                        CreatedAt: nil,
                },
	}
	return article
}

func ArticleList() []models.ArticleDetail {
        articles := []models.ArticleDetail{
                ArticleDetail("something"),
                ArticleDetail("else"),
        }
        return articles
}
