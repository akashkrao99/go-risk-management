package risks

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RisksService interface {
	CreateRisk(ctx *gin.Context, req *CreateRiskReq) (*Risk, error)
	GetRiskById(ctx *gin.Context, id string) (*Risk, error)
	GetRisks(ctx *gin.Context) []*Risk
}

type RisksServiceImplementation struct {
	repository RisksRepository
}

func NewRisksService() RisksService {
	repository := NewRisksRepository()
	return &RisksServiceImplementation{
		repository: repository,
	}
}

func (service *RisksServiceImplementation) CreateRisk(ctx *gin.Context, req *CreateRiskReq) (*Risk, error) {
	// validating request
	err := req.isValid()
	if err != nil {
		fmt.Errorf("validation failed for request : %+v", req)
		return nil, err
	}

	riskId := uuid.New().String()
	timestamp := time.Now().UnixMilli()
	newRisk := &Risk{
		Id:          riskId,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		CreatedAt:   timestamp,
	}

	createdRisk, err := service.repository.CreateRisk(ctx, newRisk)
	if err != nil {
		fmt.Errorf("error while creating risk : %+v", err)
		return nil, err
	}
	return createdRisk, nil

}

func (service *RisksServiceImplementation) GetRiskById(ctx *gin.Context, id string) (*Risk, error) {
	return service.repository.GetRiskById(ctx, id)
}

func (service *RisksServiceImplementation) GetRisks(ctx *gin.Context) []*Risk {
	return service.repository.GetRisks(ctx)
}
