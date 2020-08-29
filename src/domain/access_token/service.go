package access_token

import "github.com/Anatol-e/bookstore_oauth_api/src/domain/utils/errors"

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(id string) (*AccessToken, *errors.RestErr) {
	accessToken, err := s.repository.GetById(id)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
