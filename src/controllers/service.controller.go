package controllers

import (
	"errors"
	"fmt"
	apiErr "ka/pkg/errors"
	"ka/src/response"
	service "ka/src/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ServiceHandler struct {
	serviceService *service.ServiceService
}

func NewServiceHandler(serviceService *service.ServiceService) *ServiceHandler {
	return &ServiceHandler{serviceService}
}

func (h *ServiceHandler) GetServices(ctx *gin.Context) {
	name := ctx.Query("name")
	sort := ctx.DefaultQuery("sort", "id")
	order := ctx.DefaultQuery("order", "asc")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	services, err := h.serviceService.GetAllServices(ctx, name, sort, order, limit, offset)
	if err != nil {
		response.SendResponse(ctx, nil, apiErr.InternalError, err)
		return
	}
	response.SendResponse(ctx, services, apiErr.RequestSucess, err)
}

func (h *ServiceHandler) GetService(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.SendResponse(ctx, nil, apiErr.ErrInvalidRequest, err)
		return
	}

	service, err := h.serviceService.GetServiceByID(ctx, id)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		response.SendResponse(ctx, nil, apiErr.ErrNoServiceFound, err)
		return
	}

	if err != nil {
		response.SendResponse(ctx, nil, apiErr.InternalError, err)
		return
	}
	fmt.Println("=====================")
	fmt.Println(service)
	response.SendResponse(ctx, service, apiErr.RequestSucess, err)
}

func (h *ServiceHandler) GetServiceVersions(ctx *gin.Context) {
	serviceID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.SendResponse(ctx, nil, apiErr.ErrInvalidRequest, err)
		return
	}

	versions, err := h.serviceService.GetVersionsByServiceID(serviceID)
	if err != nil {
		response.SendResponse(ctx, nil, apiErr.ErrNoServiceFound, err)
		return
	}
	response.SendResponse(ctx, versions, apiErr.RequestSucess, err)
}
