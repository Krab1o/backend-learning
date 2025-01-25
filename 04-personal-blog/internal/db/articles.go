package db

import (
	"encoding/json"
	"log"
	"os"

	"personal-blog/internal/types"
)

func remove(s types.ArticleStorage, i uint) types.ArticleStorage {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

func openArticles() types.ArticleStorage {
	file, err := os.ReadFile(dataPath)
	articles := types.ArticleStorage{}
	if err != nil {
		log.Printf("DB: failed to open articles. %v", err)
	}
	json.Unmarshal(file, &articles)
	return articles
}

func saveID(d *data) {
	d.ID++;
	bytes, err := json.Marshal(d)
	if err != nil {
		log.Printf("DB: failed to save ID to file, %v", err)
	}
	err = os.WriteFile(idPath, bytes, os.FileMode(dataPermissions))
	if err != nil {
		log.Printf("DB: failed to save ID to file, %v", err)
	}
	
}

func saveArticles(s types.ArticleStorage) {
	bytes, err := json.Marshal(s)
	if err != nil {
		log.Printf("DB: failed to save articles to file, %v", err)
	}
	err = os.WriteFile(dataPath, bytes, os.FileMode(dataPermissions))
	if err != nil {
		log.Printf("DB: failed to save articles to file, %v", err)
	}
}

func GetArticleByID(ID uint) (*types.Article, bool) {
	articles := openArticles()
	for _, article := range articles {
		if (article.ID == ID) {
			return &article, true
		}
	}
	return &types.Article{}, false
}

func GetArticles() types.ArticleStorage {
	return openArticles()
}

func DeleteArticle(ID uint) bool {
	articles := openArticles()
	var ok bool
	for ind, article := range articles {
		if (article.ID == ID) {
			ok = true
			articles = remove(articles, uint(ind))
		}
	}
	saveArticles(articles)
	return ok
}

func CreateArticle(newArticle *types.Article) {
	articles := openArticles()
	newArticle.ID = d.ID
	saveID(&d)
	articles = append(articles, *newArticle)
	saveArticles(articles)
}

func UpdateArticle(id uint, updatedArticle *types.Article) bool {
	articles := openArticles()
	var ok bool
	for ind := 0; ind < len(articles); ind++ {
		if articles[ind].ID == id {
			ok = true
			articles = remove(articles, uint(ind))
		}
	}
	if ok {
		updatedArticle.ID = d.ID
		saveID(&d)
		articles = append(articles, *updatedArticle)
		saveArticles(articles)
	}
	return ok
}