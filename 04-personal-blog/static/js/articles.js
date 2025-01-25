'use strict';

getArticles()

async function getArticles() {
    let articlesUrl = "http://localhost:8080/api/articles"
    let articles = await fetch(articlesUrl, {
        method: "GET",
    });
    if (articles.status === 200) {
        renderArticles(articles);
    } else {
        console.log("no articles");
    }
}

async function renderArticles(response) {
    let articles = await response.json();
    console.log(articles)
    let list = document.querySelector(".articles__list");
    for (let article of articles) {
        list.appendChild(renderArticle(article));
    }
    
}

const articleRedirect = (id) => () => {
    location.href = `http://localhost:8080/articles/${id}`
}

function renderArticle(article) {
    let divArticle = document.createElement("div");
    divArticle.className = "articles__item";

    let divArticleTitle = document.createElement("div")
    divArticleTitle.className = "lower-title"
    divArticleTitle.appendChild(document.createTextNode(article.title))
    divArticle.appendChild(divArticleTitle)

    let divArticleDate = document.createElement("div")
    divArticleDate.className = "date"
    let receivedTime = new Date(article.creating_time).toDateString()
    divArticleDate.appendChild(document.createTextNode(receivedTime))
    divArticle.appendChild(divArticleDate)

    divArticle.addEventListener("click", articleRedirect(article.id))

    return divArticle;
}