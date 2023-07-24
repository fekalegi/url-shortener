package main

import (
	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
	"strings"
	"url-shortener/delivery/http/shortener"
	shortenerDomain "url-shortener/domain/shortener"
	"url-shortener/initiator"
)

func main() {
	LoadEnvVars()
	i := initiator.NewInit()

	r := i.GetGin()
	memCache := i.GetMemCache()
	api := r.Group("/api")

	r.Use(cors.Default())

	shortenerRepo := shortenerDomain.NewShortenerRepository(memCache)
	newPostService := shortenerDomain.NewShortenerService(shortenerRepo)
	shortenerController := shortener.NewShortenerController(newPostService)

	shortenerController.Route(api)

	r.Run("localhost:7000")
}

func LoadEnvVars() {
	cwd, _ := os.Getwd()
	dirString := strings.Split(cwd, "url-shortener")
	dir := strings.Join([]string{dirString[0], "url-shortener"}, "")
	AppPath := dir

	godotenv.Load(filepath.Join(AppPath, "/.env"))
}
