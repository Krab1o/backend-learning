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
    let list = document.querySelector(".articles__list")
    for (let article of articles) {
        let newArticle = document.createElement("div");
        for (let el in article) {
            
            newArticle.appendChild(document.createTextNode(article[el]));
        }
        list.appendChild(newArticle)
    }
}