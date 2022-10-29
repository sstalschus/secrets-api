package secret

import (
	"context"
	"fmt"
	"github.com/SamStalschus/secrets-api/infra/auth"
	apierr "github.com/SamStalschus/secrets-api/infra/errors"
	"github.com/SamStalschus/secrets-api/infra/log"
	"github.com/SamStalschus/secrets-api/infra/mongodb/user_repo"
	"github.com/SamStalschus/secrets-api/internal"
)

type Service struct {
	logger     log.Provider
	repository user_repo.IRepository
	apiErr     apierr.Provider
	auth       auth.Provider
}

func NewService(
	logger log.Provider,
	repository user_repo.IRepository,
	apiErr apierr.Provider,
	auth auth.Provider,
) Service {
	return Service{
		logger:     logger,
		repository: repository,
		apiErr:     apiErr,
		auth:       auth,
	}
}

func (s Service) CreateSecret(ctx context.Context, secret *internal.Secret, userID string) (apiErr *apierr.Message) {
	user, _ := s.repository.FindUserByID(ctx, userID)

	user.Secrets = append(user.Secrets, *secret)

	err := s.repository.UpdateUserByID(ctx, userID, user)
	if err != nil {
		s.logger.Info(ctx, fmt.Sprintf("Error to create one secret by user %s", userID), log.Body{})
		return s.apiErr.BadRequest("Error to create secret", err)
	}

	return apiErr
}
