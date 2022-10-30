package bookmark

import (
	"context"

	models "github.com/fernandormoraes/go-clean-architecture/domain/entities"
)

type BookmarkRepository interface {
	CreateBookmark(ctx context.Context, user *models.User, bm *models.Bookmark) error
	GetBookmarks(ctx context.Context, user *models.User) ([]*models.Bookmark, error)
	DeleteBookmark(ctx context.Context, user *models.User, id string) error
}
