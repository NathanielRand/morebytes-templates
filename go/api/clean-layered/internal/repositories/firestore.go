package repositories

import (
	"context"

	"firebase.google.com/go/v4/db"
	"github.com/NathanielRand/morebytes-templates/go/api/clean-layered/internal/models"
)

// FirestoreRepository is a repository that retrieves data from Firestore.
type FirestoreRepository struct {
	client *db.Client
}

// NewFirestoreRepository creates a new FirestoreRepository.
func NewFirestoreRepository(client *db.Client) *FirestoreRepository {
	return &FirestoreRepository{
		client: client,
	}
}

// GetImages retrieves all images from Firestore.
func (r *FirestoreRepository) GetImages(ctx context.Context) ([]*models.Image, error) {
	// TODO: Implement Firestore query to retrieve images
	return nil, nil
}

// GetImage retrieves an image from Firestore.
func (r *FirestoreRepository) AddImage(ctx context.Context, image *models.Image) error {
	// TODO: Implement Firestore query to add image
	return nil
}
