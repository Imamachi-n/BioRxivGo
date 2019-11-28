package route

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Id          int64
	Title       string
	Author      string
	Link        string
	Description string
	Published   string
	Doi         string
}

type Articles []*Article

func goHome(c *gin.Context) {

}

func GetArticlesAll(c *gin.Context) {
	// Connect DB
	db, err := connectDB()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// Close DB
	defer closeDB(db)

	// Get all articles data
	var articles Articles
	db.Find(&articles)

	// Send JSON as a response
	c.JSON(http.StatusOK, articles)
}

func PostArticle(c *gin.Context) {
	// Connect DB
	db, err := connectDB()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// Close DB
	defer closeDB(db)

	var json Article
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&json); err.Error != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
	}

	c.String(http.StatusOK, "OK")
}

func PutArticle(c *gin.Context) {
	// Connect DB
	db, err := connectDB()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// Close DB
	defer closeDB(db)

	var json Article
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get an article with a given DOI
	doi := c.Param("doi")
	var articles Articles
	db.Where("doi = ?", doi).Find(&articles)
	json.Id = articles[0].Id

	if json.Title != "" {
		articles[0].Title = json.Title
	}
	if json.Author != "" {
		articles[0].Author = json.Author
	}
	if json.Link != "" {
		articles[0].Link = json.Link
	}
	if json.Description != "" {
		articles[0].Description = json.Description
	}
	if json.Published != "" {
		articles[0].Published = json.Published
	}
	if json.Doi != "" {
		articles[0].Doi = json.Doi
	}
	db.Save(&articles[0])

	// c.String(http.StatusOK, "OK")
	c.JSON(http.StatusOK, articles[0])
}

func DeleteArticle(c *gin.Context) {
	// Connect DB
	db, err := connectDB()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// Close DB
	defer closeDB(db)

	// Delete a given article
	doi := c.Param("doi")
	var articles Articles
	db.Where("doi = ?", doi).Find(&articles)

	if len(articles) > 0 {
		db.Delete(&articles[0])
	}

	c.String(http.StatusOK, "OK")
	// c.JSON(http.StatusOK, articles[0])
}

func GetAction(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + " is " + action
	c.String(http.StatusOK, message)
}

func GetWelcome(c *gin.Context) {
	// /welcome?firstname=Jane&lastname=Doe
	firstname := c.DefaultQuery("firstname", "Guest") // If not exists, return Guest in this case
	lastname := c.Query("lastname")                   // if not exists, return "" empty string
	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}
