package main

import (
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type url struct {
	ShortenedID string `json:"shortenedID"`
	URL         string `json:"URL"`
	Timestamp   string `json:"timestamp"`
}
type reqURL struct {
	URL string `json:"id"`
}

const values = "abcdefghijkmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const baseURL = "localhost"

var urls = []url{}

func createURL(c *gin.Context) {
	// take in URL
	var newURL reqURL
	var url url
	var newURLString = newURL.URL

	if err := c.BindJSON(&newURL); err != nil {
		return
	}

	if !strings.HasPrefix(newURLString, "http://") && !strings.HasPrefix(newURLString, "https://") {
		newURLString = "https://" + newURLString
	}

	url.ShortenedID = generateNewID()
	var t = time.Now()
	url.Timestamp = t.String()
	url.URL = newURLString
	urls = append(urls, url)
	c.JSON(http.StatusOK, "http://"+baseURL+":8080/"+url.ShortenedID)
}
func deleteURL(c *gin.Context) { //TODO
	//var urlToBeDeleted url
	//look for apikey in header, if valid, continue
	// look for URL to delete, if existing in DB, delete record.
	c.JSON(http.StatusOK, "OK")
}

func generateNewID() string {
	var newString string
	var maxLength int = 7
	runes := []rune(values)
	for i := range maxLength {
		var randValue int = rand.Intn(len(runes))
		newString += string(runes[randValue])
		i += 1
	}
	return newString
}
func getAllUrls(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, urls)

}
func getURL(c *gin.Context) {
	id := c.Param("shortenedID")
	for _, url := range urls {
		if url.ShortenedID == id {
			c.Redirect(http.StatusMovedPermanently, url.URL)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "URL not found"})

}
func main() {
	router := gin.Default()
	router.POST("createURL", createURL)
	router.POST("deleteURL", deleteURL)
	router.POST("getALLURLS", getAllUrls)
	router.GET("/:shortenedID", getURL)
	//router.GET(":shortenedID", getURL)
	router.Run(":8080")
}
