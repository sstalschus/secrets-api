package secret_repo

import (
	"context"
	"github.com/SamStalschus/secrets-api/infra/mongodb"
	"github.com/SamStalschus/secrets-api/internal"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository struct {
	repository mongodb.IRepository
}

func NewRepository(
	repository mongodb.IRepository,
) Repository {
	return Repository{
		repository: repository,
	}
}

const collection = "secrets"

func (r Repository) CreateSecret(ctx context.Context, secret *internal.Secret, userID string) error {
	secret.UserID, _ = primitive.ObjectIDFromHex(userID)
	_, err := r.repository.InsertOne(ctx, collection, secret)
	return err
}

func (r Repository) FindAllByUserId(ctx context.Context, userID string) (secrets []internal.Secret) {
	objectID, _ := primitive.ObjectIDFromHex(userID)

	cursor, err := r.repository.Find(ctx, collection, bson.M{"user_id": objectID}, nil)
	if err != nil {
		return nil
	}

	err = cursor.All(ctx, &secrets)
	if err != nil {
		return nil
	}

	return secrets
}

func (r Repository) GenerateID() primitive.ObjectID {
	return primitive.NewObjectID()
}
