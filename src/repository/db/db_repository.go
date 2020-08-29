package db

import (
	"github.com/Anatol-e/bookstore_oauth_api/src/domain/access_token"
	"github.com/Anatol-e/bookstore_oauth_api/src/domain/utils/errors"
)

func NewRepository() RepositoryDb {
	return &repositoryDb{}
}

type RepositoryDb interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type repositoryDb struct {
}

func (r *repositoryDb) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServerError("database connection not implemented yet")
}
