package rest

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	repository := NewUserRepository()
	user, err := repository.LoginUser("email1234@gmail.com", "12345")

	if err != nil {
		t.Errorf("Response executed with err %s\n", err.Error)
	}
	if &user.Id == nil || user.Id == 0 {
		t.Errorf("Response cant get correct user %v\n", &user)
	}
}
