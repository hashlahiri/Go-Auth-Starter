package security

import (
	"github.com/gin-gonic/gin"
	"go-auth-starter.app/base/response"
	"go-auth-starter.app/base/utility"
)

/** Validate if user is 'CONSUMER' */
func IsUserAuthenticated(context *gin.Context) (response.AuthTokenResponse, bool) {

	isValid := false
	/** Get the userInfo from auth token */
	authResponse, _ := context.Get("userInfo")
	userInfo, _ := authResponse.(response.AuthTokenResponse)
	if(userInfo.Role == utility.Consumer) { isValid = true }

	return userInfo, isValid
}

/** Validate if user is 'ADMIN' */
func IsAdminAuthenticated(context *gin.Context) (response.AuthTokenResponse, bool) {

	isValid := false
	/** Get the userInfo from auth token */
	authResponse, _ := context.Get("userInfo")
	userInfo, _ := authResponse.(response.AuthTokenResponse)
	if(userInfo.Role == utility.Admin) { isValid = true }

	return userInfo, isValid

}