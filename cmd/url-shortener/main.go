package main

import (
	"github.com/akhilsomanvs/url-shortener/internal/config"
	"github.com/akhilsomanvs/url-shortener/internal/routes"
	"github.com/akhilsomanvs/url-shortener/internal/storage/db"
	"github.com/gin-gonic/gin"
)

func main() {
	//load config
	cfg := config.MustLoad()

	//database setup
	db := db.InitDB(cfg)
	//setup router
	server := gin.Default()
	//setup server
	routes.RegisterRoutes(server, &db)
	server.Run(cfg.Addr)
}
