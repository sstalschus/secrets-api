package user

import (
	"context"
	"fmt"

	apierr "github.com/SamStalschus/secrets-api/infra/errors"
	"github.com/SamStalschus/secrets-api/infra/hash"

	"github.com/SamStalschus/secrets-api/infra/log"
	"github.com/SamStalschus/secrets-api/infra/mongodb/user_repo"
	"github.com/SamStalschus/secrets-api/internal"
)

type Service struct {
	logger     log.Provider
	repository user_repo.IRepository
	apiErr     apierr.Provider
	auth       hash.Provider
}

func NewService(
	logger log.Provider,
	repository user_repo.IRepository,
	apiErr apierr.Provider,
	auth hash.Provider,
) Service {
	return Service{
		logger:     logger,
		repository: repository,
		apiErr:     apiErr,
		auth:       auth,
	}
}

func (s Service) CreateUser(ctx context.Context, user *internal.User) (apiErr *apierr.Message) {
	userAlreadyExists, _ := s.repository.FindUserByEmail(ctx, user.Email)
	if userAlreadyExists != nil {
		return s.apiErr.BadRequest("User already exists", fmt.Errorf(""))
	}

	user.Id = s.repository.GenerateID()

	passwordHash, err := s.auth.Encrypt(user.Password, user.Id.Hex())
	if err != nil {
		return s.apiErr.InternalServerError(err)
	}

	user.Password = string(passwordHash)

	id, err := s.repository.CreateUser(ctx, user)
	if err != nil {
		return s.apiErr.BadRequest("Error in register process", err)
	}

	s.logger.Info(ctx, fmt.Sprintf("User created: %s", id), log.Body{})

	return apiErr
}

func (s Service) GetUser(ctx context.Context, userID string) (user *internal.User, apiErr *apierr.Message) {
	user, _ = s.repository.FindUserByID(ctx, userID)

	if user == nil {
		apiErr = s.apiErr.BadRequest("User don't exists", fmt.Errorf(""))
	}
	return user, apiErr
}