package rest

import (
	"github.com/Anatol-e/bookstore_oauth_api/src/domain/users"
	"github.com/Anatol-e/bookstore_oauth_api/src/domain/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
	"time"
)

var userRestClient = rest.RequestBuilder{
	BaseURL: "localhost:8080",
	Timeout: 100 * time.Millisecond,
}

type RestUserRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type userRepository struct{}

func NewRepository() RestUserRepository {
 return &userRepository{}
}

func (u *userRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	panic("implement me")
}


