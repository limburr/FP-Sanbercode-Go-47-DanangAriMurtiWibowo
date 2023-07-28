package routes

import (
	controller "finalproj/controllers"
	middlewares "finalproj/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	routes := gin.Default()

	// set db  to gin context
	routes.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	routes.POST("/register", controller.Register)
	routes.POST("/login", controller.Login)

	// Category
	routes.GET("/categories", controller.GetAllCategory)
	routes.GET("/categories/:id", controller.GetCategoryById)
	categoryMiddlewareRoute := routes.Group("/categories")
	categoryMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	routes.POST("/", controller.CreateCategory)
	routes.PATCH("/:id", controller.UpdateCategory)
	routes.DELETE("/:id", controller.DeleteCategory)
	// routes.DELETE("/categories/:id", controller.DeleteCategory

	// Post
	routes.GET("/posts", controller.GetAllPost)
	routes.GET("/posts/:id", controller.GetPostById)
	postsMiddlewareRoute := routes.Group("/posts")
	postsMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	routes.POST("/posts", controller.CreatePost)
	routes.PATCH("/posts/:id", controller.UpdatePost)
	routes.DELETE("/posts/:id", controller.DeletePost)

	// Comment
	routes.GET("/posts/:id/comments", controller.GetcommentsByPostId)
	routes.GET("/comments/:id", controller.GetCommentById)
	routes.POST("/posts/:id/comments", controller.CreateComment)
	routes.PATCH("/comments/:id", controller.UpdateComment)
	routes.DELETE("/comments/:id", controller.DeleteComment)

	routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return routes
}
