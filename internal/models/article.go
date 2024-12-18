package models

import (
	"errors"
)

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var ArticleList = []Article{
	Article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	Article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

func GetAllArticles() []Article {
	return ArticleList
}

func GetArticleByID(id int) (*Article, error) {
	for _, article := range ArticleList {
		if article.ID == id {
			return &article, nil
		}
	}
	return nil, errors.New("Article not found!")
}
