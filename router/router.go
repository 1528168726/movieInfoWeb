package router

import (
	"databaseHW/controllers"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func InitRouter() *gin.Engine {
	server := gin.Default()
	server.SetFuncMap(template.FuncMap{
		"isDisable":   controllers.TemplateFunIsDisable,
		"add":         controllers.TemplateAdd,
		"sub":         controllers.TemplateSub,
		"checkString": controllers.TemplateFunCheckString,
	})
	server.LoadHTMLGlob("frontEnd/*.html")
	server.Static("assets", "frontEnd/assets")
	server.Static("image", "frontEnd/image")
	server.StaticFile("/favicon.ico", "frontEnd/favicon.ico")

	server.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/index")
	})
	server.GET("/actors", controllers.HandlerActors)
	server.GET("/celebrity/:id", controllers.HandlerCelebrity)
	server.GET("/directors", controllers.HandlerDirectors)
	server.GET("/index", controllers.HandlerIndex)
	server.GET("/movies", controllers.HandlerMovies)
	server.GET("/search", controllers.HandlerSearch)
	server.GET("/subject/:id", controllers.HandlerSubject)
	server.GET("/top250", controllers.HandlerTop250)
	server.GET("/Oscar/:session", controllers.HandlerOscar)
	server.GET("/Oscar", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/Oscar/92")
	})
	return server
}
