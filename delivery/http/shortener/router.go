package shortener

import (
	"github.com/gin-gonic/gin"
	"url-shortener/domain/shortener"
)

type controller struct {
	shortenerService shortener.Service
}

// NewShortenerController : Instance for register Shortener Service
func NewShortenerController(shortenerService shortener.Service) *controller {
	return &controller{shortenerService: shortenerService}
}

func (c *controller) Route(e *gin.RouterGroup) {
	v1 := e.Group("/v1")
	v1.POST("/shorten/", c.CreateShortenedURL)
	v1.GET("/shorten/:url", c.Get)
	v1.GET("/sorted_urls/", c.GetSortedURLs)
}
