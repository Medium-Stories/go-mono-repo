package web

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(cors.Default())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	return router
}
