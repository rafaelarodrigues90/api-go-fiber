package routers

import (
	"github.com/labstack/echo"
	"github.com/rafaelarodrigues90/usuarios/controllers"
)

// App é uma instancia de echo
var App *echo.Echo

func init() {
	App = echo.New()

	// homepage
	App.GET("/", controllers.Home)

	api := App.Group("/v1")

	api.GET("/read", controllers.Read)
	api.POST("/create", controllers.Create)
	api.POST("/update/:id", controllers.Update)
	api.DELETE("/delete/:id", controllers.Delete)

}
