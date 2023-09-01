package models

import "go.starlark.net/lib/time"

type SignUpUserInput struct {
	Name            string `json:"name" bson:"name"`
	Email           string `json:"email" bson:"email"`
	Password        string `json:"password" bson:"password"`
	PasswordConfirm string `json:"passwordconfirm" bson:"passwordconfirm"`
}
type SignUpUserResponse struct {
	Id        string    `json:"id" bson:"id"`
	Name      string    `json:"name" bson:"name"`
	Email     string    `json:"email" bson:"email"`
	Role      string    `json:"role" bson:"role"`
	CreatedAt time.Time `json:"createdat" bson:"createdat"`
	UpdatedAt time.Time `json:"uodatedat" bson:"updatedat"`
}

type SignInUserInput struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type SignInResponse struct {
	Status       string `json:"status" bson:"status"`
	AccessToken  string `json:"accesstoken" bson:"accesstoken"`
	RefreshToken string `json:"refreshtoken" bson:"refreshtoken"`
}
type User struct {
}
