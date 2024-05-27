package daos

import (
	"ka/pkg/db"
	"ka/src/models"
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var userDaoOnce sync.Once
var userDaoInstance *UserDao

// UserDao is a struct that represents the user dao.
type UserDao struct {
	db *gorm.DB
}

// GetUserDao is a function that returns the user dao instance.
func GetUserDao() *UserDao {
	userDaoOnce.Do(func() {
		db, err := db.InitDB()
		if err != nil {
			panic(err)
		}
		userDaoInstance = &UserDao{
			db: db,
		}
	})
	return userDaoInstance
}

// GetUser is a function that returns the user.
func (ud *UserDao) GetUser(ctx *gin.Context, tablename, userid string) (*models.User, error) {
	// Add logic here
	var user models.User
	if err := ud.db.WithContext(ctx).Table(tablename).Where("email = ?", userid).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
