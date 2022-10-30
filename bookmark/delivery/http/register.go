package http

import (
	"github.com/fernandormoraes/go-clean-architecture/bookmark"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uc bookmark.UseCase) {
	h := NewHandler(uc)

	bookmarks := router.Group("/bookmarks")
	{
		bookmarks.POST("", h.Create)
		bookmarks.GET("", h.Get)
		bookmarks.DELETE("", h.Delete)
	}
}
