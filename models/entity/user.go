package entity

import (
	"fmt"
	"log"
	"time"

	"github.com/hepiska/todo-go/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/dgrijalva/jwt-go"
)

//User struct is to handle user data
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"  json:"id"`
	Email     string             `idx:"{email},unique" json:"email" bson:"email"  binding:"required"`
	Password  string             `json:"password" binding:"required" bson:"password" `
	Name      string             `json:"name" bson:"name"`
	Address   string             `json:"address" bson:"address"`
	CreatedAt time.Time          `json:"created_at, omitempty" bson:"_created"`
	UpdatedAt time.Time          `json:"updated_at, omitempty" bson:"_modified"`
}

//GetJwtToken returns jwt token with user email claims
func (user *User) GetJwtToken() (string, error) {
	fmt.Println("jwt token email is : ", user.Email)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": string(user.Email),
	})
	log.Println(token)

	secretKey := utils.EnvVar("TOKEN_KEY")
	tokenString, err := token.SignedString([]byte(secretKey))
	return tokenString, err
}

// func init() {
// 	mogo.ModelRegistry.Register(User{})
// }
