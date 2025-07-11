// package routes

// import (
//     "echo_api/controllers"
//     "echo_api/middleware"
//     "github.com/labstack/echo/v4"
// )

// func Init() *echo.Echo {
//     e := echo.New()

//     e.POST("/register", controllers.Register)
//     e.POST("/login", controllers.Login)

//     api := e.Group("/posts")
//     api.Use(middleware.JWTMiddleware)
//     api.POST("", controllers.CreatePost)
//     api.GET("", controllers.GetPosts)
//     api.PUT("/:id", controllers.UpdatePost)
//     api.DELETE("/:id", controllers.DeletePost)

//     return e
// }

package routes

import (
	"echo_api/controllers"
	"echo_api/middleware"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)
	e.POST("/refresh", controllers.RefreshToken)

	api := e.Group("/posts")
	api.Use(middleware.JWTMiddleware)
	api.POST("", controllers.CreatePost)
	api.GET("", controllers.GetPosts)
	api.PUT("/:id", controllers.UpdatePost)
	api.DELETE("/:id", controllers.DeletePost)

	return e
}
