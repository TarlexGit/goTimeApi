package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/time")
	{
		api.GET("/now", h.getNow)
		api.GET("/string", h.getString)
		api.GET("/add", h.getAdd)
		api.POST("/correct", h.postCorrect)
	}
	return router
}
