package newsapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Article struct {
	Author      string
	Title       string
	Description string
	URL         string
	PublishedAt time.Time `json:publishedAt`
	Content     string
	UrlToImage  string `json:urlToImage`
}

type news struct {
	Articles     []Article
	TotalResults int32 `json:totalResults`
	Status       string
}

func reformat_date(date time.Time) string {
	return date.Format("2006-01-02")
}

// Search news terms on the specific date
func SearchForDate(terms string, date time.Time) ([]Article, error) {
	url := fmt.Sprintf("https://newsapi.org/v2/everything?q=%s&from=%s", terms, reformat_date(date))
	return search(url)
}

// Search news terms on a specific date and domain
func SearchForDateAndDomain(terms string, date time.Time, domain string) ([]Article, error) {
	url := fmt.Sprintf("https://newsapi.org/v2/everything?q=%s&from=%s&domain=%s", terms, reformat_date(date), domain)
	return search(url)
}

// Search news terms on a specific domain
func SearchForDomain(terms string, domain string) ([]Article, error) {
	url := fmt.Sprintf("https://newsapi.org/v2/everything?q=%s&domain=%s", terms, domain)
	return search(url)
}

// Search news terms
func Search(terms string) ([]Article, error) {
	url := fmt.Sprintf("https://newsapi.org/v2/everything?q=%s", terms)
	return search(url)
}

func search(baseURL string) ([]Article, error) {
	url := fmt.Sprintf("%s&apiKey=%s", baseURL, token)
	resp, _ := http.Get(url)
	if resp.StatusCode >= 204 {
		body, _ := ioutil.ReadAll(resp.Body)
		err := fmt.Errorf("Failed to fetch news:\n%s", body)
		return nil, err
	}

	news := news{}
	defer resp.Body.Close()
	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return news.Articles, e
	}
	json.Unmarshal(body, &news)
	return news.Articles, nil
}
