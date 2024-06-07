package domain

import (
	"context"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"go-auth-starter.app/base/utility"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection * mongo.Collection

func SetUserCollection(client *mongo.Client) {
	userCollection = client.Database("go-auth-starter").Collection("users")
}

type User struct {
	ID 			primitive.ObjectID `bson:"_id,omitempty"`
	Username 	string `bson:"username"`
	Email		string `bson:"email"`
	Contact		string `bson:"phone"`
	Password	string `bson:"password" validate:"required"`
	CreatedAt	time.Time `bson:"createdAt"`
	ModifiedAt	time.Time `bson:"modifiedAt"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// Validate function to check struct fields
func (u User) Validate() error {
	return validate.Struct(u)
}

//var users = []User{}

/** Get user by id */
func GetUserById(ctx context.Context, id string) (User, error) {

	var user User

	objectID, _ := primitive.ObjectIDFromHex(id)

	err := userCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, err
		}
		return user, err
	}

	return user, nil
}

/** Get user by username */
func GetUserByUsername(ctx context.Context, username string) (User, error) {

	var user User

	err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, err
		}
		return user, err
	}

	return user, nil
}

/** save user */
func (userToSave *User) SaveUser(ctx context.Context) (*mongo.InsertOneResult, error) {

	if err := userToSave.Validate(); err != nil {
		return nil, err
	}

	/** 'Username' - converted */
	if(userToSave.Email != "") {
		userToSave.Username = userToSave.Email
	} else {
		userToSave.Username = userToSave.Contact
	}

	/** 'Password' - management */
	hashedPassword, err := utility.HashPassword(userToSave.Password)
	if(err != nil) {
		return nil, err
	}

	userToSave.Password = hashedPassword

	userToSave.CreatedAt = time.Now().Local()
	userToSave.ModifiedAt = time.Now().Local()

	userSaved, err := userCollection.InsertOne(ctx, userToSave)
	
	return userSaved, err
}

/** validate user and generate token */
func (userToValidate *User) ValidateUserAndGenerateToken(ctx context.Context) (string, error) {

	/** Fetch existing user by username */
	userFetched, err := GetUserByUsername(ctx, userToValidate.Username)
	if(err != nil) {
		return "", err
	}

	passwordMatchResult := utility.CompareHashedPasswords(userFetched.Password, userToValidate.Password)
	if passwordMatchResult {

		// generate the token, user matched
		token, err := utility.GenerateToken(userFetched.Username, userFetched.ID.Hex())
		return token, err
	} else {
		
		return "credentials did not match", errors.New("credentials did not match")
	}
	
}