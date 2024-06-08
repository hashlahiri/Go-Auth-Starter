package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-auth-starter.app/base/domain"
	"go-auth-starter.app/base/response"
	"go-auth-starter.app/base/security"
)

// ADMIN
/** Get user by id */
func getUserById(context *gin.Context) {

	/** Check if 'ADMIN' authentication or not */
	_, isUser := security.IsAdminAuthenticated(context)
	if(!isUser) {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Authentication needed"})
		return
	}

	/** fetch by path variable 'id' */
	id := context.Param("id")
	if id == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "'id' is missing from input param"})
		return
	}

	/** Fetch userInfo */
	userInfoFetched, err := domain.GetUserById(context, id) // get user by id
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "no user found by given 'id'"})
		return
	}

	/** return response */
	context.JSON(http.StatusOK, userInfoFetched)
}

/** Get user identifying from token passed */
func getMyUser(context *gin.Context) {

	// /** Get the userInfo from auth token */
	// authResponse,_ := context.Get("userInfo")
	// userInfo,_ := authResponse.(response.AuthTokenResponse)
	userInfo, isUser := security.IsUserAuthenticated(context)
	if(!isUser) {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Authentication needed"})
		return
	}

	/** Fetch associated user */
	userInfoFetched, err := domain.GetUserById(context, userInfo.Id) // get user by id
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "no user found by given 'id'"})
		return
	}

	/** convert to consumer response */
	var response = convertUserInfoToConsumerResponse(userInfoFetched)

	context.JSON(http.StatusOK, response)
}

/** Convert user object to consumer user response */
func convertUserInfoToConsumerResponse(userInfo domain.User) (response.UserConsumerResponse) {

	/** declare object */
	consumerUserResponse := response.UserConsumerResponse{}

	/** assign values to object */
	consumerUserResponse.ID = userInfo.ID
	consumerUserResponse.Username = userInfo.Username
	consumerUserResponse.Email = userInfo.Email
	consumerUserResponse.Contact = userInfo.Contact

	return consumerUserResponse
}

// ADMIN
/** Get user by username */
func getUserByUsername(context *gin.Context) {

	/** Check if 'ADMIN' authentication or not */
	_, isUser := security.IsAdminAuthenticated(context)
	if(!isUser) {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Authentication needed"})
		return
	}
	
	/** extract from path */
	username := context.Param("username")
	if username == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "'username' is missing from input param"})
		return
	}

	/** Fetch associated user */
	userInfoFetched, err := domain.GetUserByUsername(context, username) // get user by id
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "no user found by given 'username'"})
		return
	}

	context.JSON(http.StatusOK, userInfoFetched)
}

/** Register a User */
func registerUser(context *gin.Context) {

	/** capture from body passed */
	var userToSave domain.User
	err := context.ShouldBindJSON(&userToSave)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	/** 'email' / 'contact' - one of them MUST be present */
	if(userToSave.Email == "" && userToSave.Contact == "") {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Both 'email' & 'contact' cannot be NULL"})
		return
	}

	response, err := userToSave.SaveUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not save user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created!", "user": response})
}

// AUTH GENERATOR
/** Generates the associated token for user */
func generateToken(context *gin.Context) {

	var userToValidate domain.User
	err := context.ShouldBindJSON(&userToValidate)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	/** Validation - 'username' & 'password' */
	if( (userToValidate.Username == "") || (userToValidate.Password == "") ) {

		context.JSON(http.StatusBadRequest, gin.H{"message": "'username' & 'password' is compulsory"})
		return
	}

	response, err := userToValidate.ValidateUserAndGenerateToken(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid credentials"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"token": response})

}