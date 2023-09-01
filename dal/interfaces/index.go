package interfaces

import "project-01/dal/models"

type IUser interface {
	SignIn(signin *models.SignInUserInput) *models.SignInResponse
	SignUp(signup models.SignUpUserInput) *models.SignUpUserInput
}
