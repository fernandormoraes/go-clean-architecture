package repositories

import (
	"context"

	"github.com/fernandormoraes/go-clean-architecture/data/dto"
	"github.com/fernandormoraes/go-clean-architecture/domain/entities"
)

type BookmarkRepository interface {
	CreateBookmark(ctx context.Context, user *entities.User, bookMark *dto.CreateBookmarkDTO) error
	GetBookmarks(ctx context.Context, user *entities.User) ([]*dto.ReadBookmarkDTO, error)
}
