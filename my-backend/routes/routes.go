package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadfikri13/ecommerce-app/my-backend/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) { // Routing for User and Admin
	incomingRoutes.POST("/users/signup", controllers.Signup())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addProduct", controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/productView", controllers.SearchProduct())
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())

}
