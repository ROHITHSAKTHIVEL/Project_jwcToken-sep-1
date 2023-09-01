package services

import (
	"context"
	"project-01/dal/interfaces"
	"project-01/dal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

// SignIn implements interfaces.IUser.
func (*UserService) SignIn(signin *models.SignInUserInput) *models.SignInResponse {
	panic("unimplemented")
}

// SignUp implements interfaces.IUser.
func (*UserService) SignUp(signup models.SignUpUserInput) *models.SignUpUserInput {
	panic("unimplemented")
}

func UserInit(collection *mongo.Collection, ctx context.Context) interfaces.IUser {
	return &UserService{collection, ctx}
}

func (p *UserService) CreateCustomer(user *models.SignInUserInput) error {

	_, err := p.usercollection.InsertOne(p.ctx, &user)

	if err != nil {
		return err
	}
	return nil

}
