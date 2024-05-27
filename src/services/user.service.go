package services

import (
	"ka/constants"
	"ka/src/daos"
	"sync"

	apiErr "ka/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

var userServiceOnce sync.Once
var userServiceInstance *UserService

// UserService is a struct that represents the user service.
type UserService struct {
	userDao *daos.UserDao
}

func (us *UserService) GetUserDao() *daos.UserDao {
	return us.userDao
}

// GetUserService is a function that returns the user service instance.
func GetUserService() *UserService {
	userServiceOnce.Do(func() {
		userServiceInstance = &UserService{
			userDao: daos.GetUserDao(),
		}
	})
	return userServiceInstance
}

// authenticate user
func (us *UserService) AuthenticateUser(ctx *gin.Context) (apiErr.Error, error) {
	email, password, found := ctx.Request.BasicAuth()
	if !found {
		return apiErr.UserNotAuth, errors.WithStack(apiErr.ErrUserNotFound)
	}

	user, err := us.GetUserDao().GetUser(ctx, constants.T_AUTH, email)
	if err != nil {
		return apiErr.InternalError, errors.WithStack(err)
	}

	if user.GetEmail() != email {
		return apiErr.UserNotAuth, errors.WithStack(apiErr.ErrWrongUserId)
	}

	if user.GetPassword() != password {
		return apiErr.UserNotAuth, errors.WithStack(apiErr.ErrWrongPassword)
	}
	return apiErr.GetDefaultError(), nil
}
