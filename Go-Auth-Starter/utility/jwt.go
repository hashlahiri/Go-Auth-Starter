package utility

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go-auth-starter.app/base/response"
)

const secretKey = "wakandaForever"

func GenerateToken(username string, userId string, userRole string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username" : username,
		"userId" : userId,
		"role" : userRole,
		"exp" : time.Now().Add(time.Hour * 2).Unix(),
	})
	
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (response.AuthTokenResponse, error) {

	response := response.AuthTokenResponse{}

	/** Check if passed in 'Bearer Token' */
	if strings.HasPrefix(token, "Bearer ") {
		token = token[7:]
	}
	log.Println("Printing upto here")

	/** Check token signing method */
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC) // Check token type
		if(!ok) {
			return nil, errors.New("unexpected signing method") // return error response
		}
		return []byte(secretKey), nil
	})
	log.Println("Still here")

	if(err != nil) {
		return response, errors.New("could not parse Token")
	}
	log.Println("3")
	tokenIsValid := parsedToken.Valid
	if(!tokenIsValid) {
		return response, errors.New("invalid Token")
	}
	log.Println("4")
	/** Type check */
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if(!ok) {
		return response, errors.New("Invalid token claims!")
	}
	log.Println("5")
	response.Username = claims["username"].(string)
	response.Id = claims["userId"].(string)
	response.ExpiresIn = time.Unix(int64(claims["exp"].(float64)), 0)
	response.Role = claims["role"].(string)
	log.Println("6")
	return response, nil
}