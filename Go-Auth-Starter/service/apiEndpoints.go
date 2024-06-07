package service

import (
	"github.com/gin-gonic/gin"
	"go-auth-starter.app/base/middleware"
)

const BasePath = "/api"

func ApiEndpoints(server *gin.Engine) {

	/** Authenticated Routes */
	authenticated := server.Group(BasePath)
	authenticated.Use(middleware.Authenticate)

	/** Endpoints defined */
	userEndpoint := "/user"

	// User endpoints
	authenticated.GET(userEndpoint+"/get/my/", getMyUser)
	authenticated.GET(userEndpoint+"/get/:id", getUserById) // ADMIN
	authenticated.GET(userEndpoint+"/get/username/:username", getUserByUsername) // ADMIN

	// Other endpoints
	server.POST("/token", generateToken)
	server.POST("/register", registerUser)
	
}