package response

import(
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserConsumerResponse struct {
	ID 			primitive.ObjectID `bson:"_id,omitempty"`
	Username 	string `bson:"username"`
	Email		string `bson:"email"`
	Contact		string `bson:"phone"`
}