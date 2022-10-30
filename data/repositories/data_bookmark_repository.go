package datarepository

import (
	"context"

	externalbookmarkdatasource "github.com/fernandormoraes/go-clean-architecture/data/datasources/external"
	"github.com/fernandormoraes/go-clean-architecture/data/dto"
	"github.com/fernandormoraes/go-clean-architecture/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DataBookmarkRepository struct {
	externalBookmark *externalbookmarkdatasource.DbBookmarkDatasource
}

func NewBookmarkRepository(externalBookmark *externalbookmarkdatasource.DbBookmarkDatasource) *DataBookmarkRepository {
	return &DataBookmarkRepository{
		externalBookmark: externalBookmark,
	}
}

func (r DataBookmarkRepository) CreateBookmark(ctx context.Context, user *entities.User, bookMark *dto.CreateBookmarkDTO) error {
	bookMarkEntity := createBookMarkToModel(bookMark)

	_, err := r.externalBookmark.CreateBookmark(ctx, user, bookMarkEntity)

	if err != nil {
		return err
	}

	return nil
}

func (r DataBookmarkRepository) GetBookmarks(ctx context.Context, user *entities.User) ([]*dto.ReadBookmarkDTO, error) {
	cursor, error := r.externalBookmark.GetBookmarks(ctx, user)

	if error != nil {
		return nil, error
	}

	out := make([]*entities.Bookmark, 0)

	for cursor.Next(ctx) {
		user := new(entities.Bookmark)
		error := cursor.Decode(user)
		if error != nil {
			return nil, error
		}

		out = append(out, user)
	}
	if error := cursor.Err(); error != nil {
		return nil, error
	}

	return toBookmarks(out), nil
}

func createBookMarkToModel(b *dto.CreateBookmarkDTO) *entities.Bookmark {
	uid, _ := primitive.ObjectIDFromHex(b.UserID)

	return &entities.Bookmark{
		UserID: uid,
		URL:    b.URL,
		Title:  b.Title,
	}
}

func toBookmark(b *entities.Bookmark) *dto.ReadBookmarkDTO {
	return &dto.ReadBookmarkDTO{
		ID:     b.ID.Hex(),
		UserID: b.UserID.Hex(),
		URL:    b.URL,
		Title:  b.Title,
	}
}

func toBookmarks(bs []*entities.Bookmark) []*dto.ReadBookmarkDTO {
	out := make([]*dto.ReadBookmarkDTO, len(bs))

	for i, b := range bs {
		out[i] = toBookmark(b)
	}

	return out
}
