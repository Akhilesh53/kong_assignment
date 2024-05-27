package daos

import (
	"ka/constants"
	"ka/pkg/db"
	"ka/src/models"
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var postgresRepoOnce sync.Once
var postgresRepo *ServiceRepository

type ServiceRepository struct {
	db *gorm.DB
}

func NewServiceRepository() *ServiceRepository {
	postgresRepoOnce.Do(func() {
		db, err := db.InitDB()
		if err != nil {
			panic(err)
		}
		postgresRepo = &ServiceRepository{db}
	})
	return postgresRepo

}

func (sr *ServiceRepository) GetAllServices(ctx *gin.Context, name string, sort string, order string, limit int, offset int) ([]models.Service, error) {
	var services []models.Service
	query := sr.db.WithContext(ctx).Table(constants.T_SERVICES)

	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	query = query.Order(sort + " " + order).Offset(offset).Limit(limit).Preload("Versions").Find(&services)
	return services, query.Error
}

func (sr *ServiceRepository) GetServiceByID(ctx *gin.Context, id int) (*models.Service, error) {
	var service models.Service
	if err := sr.db.WithContext(ctx).Table(constants.T_SERVICES).Where("id = ?", id).Error; err != nil {
		return nil, err
	}
	return &service, nil
}

func (r *ServiceRepository) GetVersionsByServiceID(serviceID int) ([]models.Version, error) {
	var versions []models.Version
	if err := r.db.Where("service_id = ?", serviceID).Find(&versions).Error; err != nil {
		return nil, err
	}
	return versions, nil
}
