package router

import (
	"database/sql"
	"quiz-pekan3-damar/controller"
	"quiz-pekan3-damar/middleware"

	"github.com/gin-gonic/gin"
)

func StartServer(db *sql.DB) *gin.Engine {
	setRouter := gin.Default()

	setRouter.POST("/api/users/register", controller.TambahUser)
	setRouter.POST("/api/users/login", controller.LoginUser)

	setJWTAPI := setRouter.Group("/api")
	setJWTAPI.Use(middleware.MiddlewareJWT())
	{
		setJWTAPI.GET("/categories", controller.GetKategori)
		setJWTAPI.GET("/categories/:id", controller.GetKategoriID)
		setJWTAPI.POST("/categories", controller.TambahKategori)
		setJWTAPI.PUT("/categories/:id", controller.UpdateKategori)
		setJWTAPI.DELETE("/categories/:id", controller.HapusKategori)

		setJWTAPI.GET("/categories/:id/books", controller.GetBukuBerdasarkanKategori)

		setJWTAPI.GET("/books", controller.GetBuku)
		setJWTAPI.GET("/books/:id", controller.GetBukuID)
		setJWTAPI.POST("/books", controller.TambahBuku)
		setJWTAPI.DELETE("/books/:id", controller.HapusBuku)
		setJWTAPI.PUT("/books/:id", controller.UpdateBuku)

	}

	return setRouter

}
