package externaldatasource

import (
	"context"
	"fmt"

	"github.com/fernandormoraes/go-clean-architecture/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DbBookmarkDatasource struct {
	db *mongo.Collection
}

func NewDbBookmarkDatasource(db *mongo.Database, collection string) *DbBookmarkDatasource {
	return &DbBookmarkDatasource{
		db: db.Collection(collection),
	}
}

func (r DbBookmarkDatasource) CreateBookmark(ctx context.Context, user *entities.User, bookMark *entities.Bookmark) (*mongo.InsertOneResult, error) {
	res, err := r.db.InsertOne(ctx, bookMark)
	fmt.Print(err)
	return res, err
}

func (r DbBookmarkDatasource) GetBookmarks(ctx context.Context, user *entities.User) (cur *mongo.Cursor, err error) {
	return r.db.Find(ctx, bson.M{})
}

func (r DbBookmarkDatasource) DeleteBookmark(ctx context.Context, user *entities.User, id string) (*mongo.DeleteResult, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	uID, _ := primitive.ObjectIDFromHex(user.ID)

	return r.db.DeleteOne(ctx, bson.M{"_id": objID, "userId": uID})
}
