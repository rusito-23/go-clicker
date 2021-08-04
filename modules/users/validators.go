package users

import (
	"github.com/gin-gonic/gin"
	"go-clicker/modules/common"
)

// UserModelValidator -
// Contains the validation logic for a user register request
type UserModelValidator struct {
	User struct {
		Username  string `form:"username" json:"username" binding:"required,alphanum,min=4,max=255"`
		FirstName string `form:"first_name" json:"first_name" binding:"required,min=2,max=100"`
		LastName  string `form:"last_name" json:"last_name" binding:"required,min=2,max=100"`
		Email     string `form:"email" json:"email" binding:"required,email"`
		Password  string `form:"password" json:"password" binding:"required,min=8,max=255"`
		Bio       string `form:"bio" json:"bio" binding:"max=1024"`
	} `json:"user"`
}

// Bind -
// Validates the request and creates a new user model
func (validator *UserModelValidator) Bind(c *gin.Context) (UserModel, error) {
	user := UserModel{}
	err := common.Bind(c, validator)
	if err != nil {
		return user, err
	}

	user = UserModel{
		Username:  validator.User.Username,
		FirstName: validator.User.FirstName,
		LastName:  validator.User.LastName,
		Email:     validator.User.Email,
		Password:  validator.User.Password,
		Bio:       validator.User.Bio,
	}

	return user, nil
}
