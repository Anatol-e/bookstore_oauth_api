package rest

import (
	"bytes"
	"encoding/json"
	"github.com/Anatol-e/bookstore_oauth_api/src/domain/users"
	"github.com/Anatol-e/bookstore_oauth_api/src/domain/utils/errors"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	BaseUrl = "http://localhost:8080"
	Timeout = 100 * time.Millisecond
)

type RestUserRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type userRepository struct{}

func NewUserRepository() RestUserRepository {
	return &userRepository{}
}

func (u *userRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	request := users.LoginRequest{
		Email:    email,
		Password: password,
	}
	byteRequest, _ := json.Marshal(request)
	response, err := http.Post(BaseUrl+"/users/login", "application/json", bytes.NewBuffer(byteRequest))
	if err != nil {
		return nil, errors.NewInternalServerError("invalid restclient response when trying to login user")
	}
	defer response.Body.Close()
	bytesBody, _ := ioutil.ReadAll(response.Body)
	if response.StatusCode > 299 {
		var restErr errors.RestErr
		err := json.Unmarshal(bytesBody, &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("invalid error interface hen trying to login user")
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(bytesBody, &user); err != nil {
		return nil, errors.NewInternalServerError("error when trying unmarshal users response")
	}
	return &user, nil
}
