package response

import "time"

type AuthTokenResponse struct {
	Username  string    `json:"username"`
	Id        string    `json:"userId"`
	ExpiresIn time.Time `json:"exp"`
}