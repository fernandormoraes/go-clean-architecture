package http

import (
	"net/http"

	"github.com/fernandormoraes/go-clean-architecture/data/dto"
	"github.com/fernandormoraes/go-clean-architecture/domain/entities"
	"github.com/fernandormoraes/go-clean-architecture/domain/usecases"
	"github.com/gin-gonic/gin"
)

type Bookmark struct {
	ID    string `json:"id"`
	URL   string `json:"url"`
	Title string `json:"title"`
}

type Handler struct {
	useCase usecases.BookmarkUseCase
}

func NewHandler(useCase usecases.BookmarkUseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

type createInput struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}

func (h *Handler) Create(c *gin.Context) {
	inp := new(createInput)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := &entities.User{
		ID:       "1",
		Username: "Fernando",
		Password: "Teste",
	}

	if err := h.useCase.CreateBookmark(c.Request.Context(), user, inp.URL, inp.Title); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

type getResponse struct {
	Bookmarks []*dto.ReadBookmarkDTO `json:"bookmarks"`
}

func (h *Handler) Get(c *gin.Context) {
	user := &entities.User{
		ID:       "1",
		Username: "Fernando",
		Password: "Teste",
	}

	bms, err := h.useCase.GetBookmarks(c.Request.Context(), user)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &getResponse{
		Bookmarks: bms,
	})
}
