package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.get)
			lists.DELETE("/:id")

			items := lists.Group(":id/items")
			{
				items.POST("/")
				items.GET("/")
				items.GET("/:item_id")
				items.PUT("/:item_id")
				items.DELETE("/:item_id")

			}
		}
	}
	return router
}
