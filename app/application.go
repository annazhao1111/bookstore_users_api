package app

import (
	"github.com/annazhao/bookstore_users_api/logger"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func StartApplication() {
	mapUrls()

	// logger.Log.Info("about to start the application...")
	logger.Info("about to start the application...")
	router.Run(":8080")
}
