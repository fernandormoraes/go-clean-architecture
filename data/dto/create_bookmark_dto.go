package dto

type CreateBookmarkDTO struct {
	UserID string `bson:"userId"`
	URL    string `bson:"url"`
	Title  string `bson:"title"`
}
