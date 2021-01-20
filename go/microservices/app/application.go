package app

import (
	"github.com/gin-gonic/gin" // presentation layer
)

var (
	router = gin.Default()
)

func StatApplication() {
	mapUrls()
	router.Run(":8000")
}
