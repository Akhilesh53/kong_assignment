package services

import (
	repository "ka/src/daos"
	"ka/src/models"

	"github.com/gin-gonic/gin"
)

type ServiceService struct {
	repo *repository.ServiceRepository
}

func NewServiceService(repo *repository.ServiceRepository) *ServiceService {
	return &ServiceService{repo}
}

func (s *ServiceService) GetAllServices(ctx *gin.Context, name string, sort string, order string, limit int, offset int) ([]models.Service, error) {
	return s.repo.GetAllServices(ctx, name, sort, order, limit, offset)
}

func (s *ServiceService) GetServiceByID(ctx *gin.Context, id int) (*models.Service, error) {
	return s.repo.GetServiceByID(ctx, id)
}

func (s *ServiceService) GetVersionsByServiceID(serviceID int) ([]models.Version, error) {
	return s.repo.GetVersionsByServiceID(serviceID)
}
