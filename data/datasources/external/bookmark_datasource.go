package externaldatasource

import (
	"context"

	"github.com/fernandormoraes/go-clean-architecture/data/dto"
	"github.com/fernandormoraes/go-clean-architecture/domain/entities"
)

type BookmarkRepository interface {
	CreateBookmark(ctx context.Context, user *entities.User, bm *dto.CreateBookmarkDTO) error
	GetBookmarks(ctx context.Context, user *entities.User) ([]*entities.Bookmark, error)
	DeleteBookmark(ctx context.Context, user *entities.User, id string) error
}
