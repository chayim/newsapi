package newsapi_test

import (
	"newsapi"
	"testing"
	"time"
)

func validate(item newsapi.Article, t *testing.T) {
	if item.Title == "" {
		t.Error("No title retrieved for article")
	}
	if item.Description == "" {
		t.Error("No description retrieved for article")
	}
	if item.URL == "" {
		t.Error("No URL retrieved for article")
	}
	if item.Content == "" {
		t.Error("No content for article")
	}
	if item.PublishedAt.Year() < 2010 {
		t.Error("bad publication year for article")
	}
}

func TestGetNews(t *testing.T) {
	articles, err := newsapi.Search("Google")
	if err != nil {
		t.Error(err)
		return
	}
	if len(articles) == 0 {
		t.Error("No articles found.")
		return
	}
	for _, item := range articles {
		validate(item, t)
	}
}

func TestGetNewsForDate(t *testing.T) {
	d := time.Now().AddDate(0, 0, -20).Format("2006-01-02")
	date, _ := time.Parse("2006-01-02", d)
	articles, err := newsapi.SearchForDate("Google", date)
	if err != nil {
		t.Error(err)
		return
	}

	if len(articles) == 0 {
		t.Error("No articles found.")
		return
	}
	for _, item := range articles {
		validate(item, t)
	}
}

func TestGetNewsForDateAndDomain(t *testing.T) {
	d := time.Now().AddDate(0, 0, -20).Format("2006-01-02")
	date, _ := time.Parse("2006-01-02", d)
	articles, err := newsapi.SearchForDateAndDomain("Google", date, "nytimes.com")
	if err != nil {
		t.Error(err)
		return
	}

	if len(articles) == 0 {
		t.Error("No articles found.")
		return
	}
	for _, item := range articles {
		validate(item, t)
	}
}

func TestGetNewsForDomain(t *testing.T) {
	articles, err := newsapi.SearchForDomain("Google", "nytimes.com")
	if err != nil {
		t.Error(err)
		return
	}

	if len(articles) == 0 {
		t.Error("No articles found.")
		return
	}
	for _, item := range articles {
		validate(item, t)
	}
}
