package models

import "testing"

func TestGetAllArticles(t *testing.T) {
	aList := GetAllArticles()

	if len(aList) != len(ArticleList) {
		t.Fail()
	}

	for i, v := range aList {
		if v.Content != ArticleList[i].Content || v.ID != ArticleList[i].ID || v.Title != ArticleList[i].Title {
			t.Fail()
			break
		}
	}

}
