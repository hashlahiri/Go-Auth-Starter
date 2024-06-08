package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-auth-starter.app/base/utility"
)

/** Authenticates an incoming request to server */
func Authenticate(context *gin.Context) {
	/**Check if contains 'Authorization'*/
	token := context.Request.Header.Get("Authorization")
	if(token == "") {

		// makes sure to abort and not execute any other apis
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization needed"})
		return
	}

	/** Check valid token */
	userInfo, err := utility.VerifyToken(token)
	if(err != nil) {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization needed"})
		return
	}

	context.Set("userInfo", userInfo)

	/** makes sure that the next request handler executes properly */
	context.Next()

}