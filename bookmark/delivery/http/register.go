package http

import (
	"github.com/fernandormoraes/go-clean-architecture/domain/usecases"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uc usecases.BookmarkUseCase) {
	h := NewHandler(uc)

	bookmarks := router.Group("/bookmarks")
	{
		bookmarks.POST("", h.Create)
		bookmarks.GET("", h.Get)
	}
}
