package handlers

import (
	"errors"
	"net/http"

	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/entity"
	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/repository"

	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	repo *repository.URLRepository
}

func NewURLHandler(repo *repository.URLRepository) URLHandler {
	return URLHandler{
		repo: repo,
	}
}

func (h *URLHandler) Get(c *gin.Context) {
	// TODO: answer here
	url, err := h.repo.Get(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"url": url,
	})

}

func (h *URLHandler) Create(c *gin.Context) {
	// TODO: answer here
	var url entity.URL
	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("invalid url"),
		})
		return
	}
	urls, err := h.repo.Create(url.LongURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "URL created",
		"data":    urls,
	})

}

func (h *URLHandler) CreateCustom(c *gin.Context) {
	// TODO: answer here
	var url entity.URL
	if method := c.Request.Method; method == "GET" {
		c.Redirect(http.StatusFound, "https://pawgrammers.com")
	}
	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	urls, err := h.repo.CreateCustom(url.LongURL, url.ShortURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "URL created",
		"data":    urls,
	})
}
