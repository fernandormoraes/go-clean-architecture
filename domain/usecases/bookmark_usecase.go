package usecases

import (
	"context"

	"github.com/fernandormoraes/go-clean-architecture/data/dto"
	"github.com/fernandormoraes/go-clean-architecture/domain/entities"
	"github.com/fernandormoraes/go-clean-architecture/domain/repositories"
)

type BookmarkUseCase struct {
	bookmarkRepo repositories.BookmarkRepository
}

func NewBookmarkUseCase(bookmarkRepo repositories.BookmarkRepository) *BookmarkUseCase {
	return &BookmarkUseCase{
		bookmarkRepo: bookmarkRepo,
	}
}

func (b BookmarkUseCase) CreateBookmark(ctx context.Context, user *entities.User, url, title string) error {
	bm := &dto.CreateBookmarkDTO{
		URL:   url,
		Title: title,
	}

	return b.bookmarkRepo.CreateBookmark(ctx, user, bm)
}

func (b BookmarkUseCase) GetBookmarks(ctx context.Context, user *entities.User) ([]*dto.ReadBookmarkDTO, error) {
	return b.bookmarkRepo.GetBookmarks(ctx, user)
}
