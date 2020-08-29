package app

import (
	"github.com/Anatol-e/bookstore_oauth_api/src/domain/access_token"
	"github.com/Anatol-e/bookstore_oauth_api/src/http"
	"github.com/Anatol-e/bookstore_oauth_api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func StartApplication() {
	atService := http.NewHandler(access_token.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atService.GetById)
	router.Run(":8080")
}
