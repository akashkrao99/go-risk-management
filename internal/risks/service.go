package risks

import (
	"fmt"
	"net/http"
	"time"

	myhttp "github.com/akashkrao99/go-sample-http/pkg/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RisksService interface {
	CreateRisk(ctx *gin.Context, req *CreateRiskReq) (*CreateRiskRes, error)
	GetRiskById(ctx *gin.Context, id string) (*GetRiskByIdRes, error)
	GetRisks(ctx *gin.Context) (*GetRisksRes, error)
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

func NewRisksServiceImplementation(repo RisksRepository) *RisksServiceImplementation {
    return &RisksServiceImplementation{repository: repo}
}

func (service *RisksServiceImplementation) CreateRisk(ctx *gin.Context, req *CreateRiskReq) (*CreateRiskRes, error) {

	// validating request
	err := req.IsValid()
	if err != nil {
		fmt.Errorf("validation failed for request : %+v", req)
		return &CreateRiskRes{
			ResponseJson: myhttp.ResponseJson{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
			},
		}, err
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

	// basic data sanitization can be performed here
	newRisk.sanitize()

	err = newRisk.isValid()
	if err != nil {
		fmt.Errorf("model validation failed : %v", err.Error())
		return &CreateRiskRes{
			ResponseJson: myhttp.ResponseJson{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
			},
		}, err
	}

	createdRisk, err := service.repository.CreateRisk(ctx, newRisk)
	if err != nil {
		fmt.Errorf("error while creating risk : %+v", err)
		return &CreateRiskRes{
			ResponseJson: myhttp.ResponseJson{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			},
		}, err
	}

	return &CreateRiskRes{
		ResponseJson: myhttp.ResponseJson{
			StatusCode: http.StatusCreated,
			Message:    http.StatusText(http.StatusOK),
		},
		Risk: createdRisk,
	}, nil

}

func (service *RisksServiceImplementation) GetRiskById(ctx *gin.Context, id string) (*GetRiskByIdRes, error) {
	risk, err := service.repository.GetRiskById(ctx, id)
	if err != nil {
		return &GetRiskByIdRes{
			ResponseJson: myhttp.ResponseJson{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			},
		}, err
	}

	return &GetRiskByIdRes{
		ResponseJson: myhttp.ResponseJson{
			StatusCode: http.StatusOK,
			Message:    http.StatusText(http.StatusOK),
		},
		Risk: risk,
	}, nil
}

func (service *RisksServiceImplementation) GetRisks(ctx *gin.Context) (*GetRisksRes, error) {
	risks, err := service.repository.GetRisks(ctx)
	if err != nil {
		return &GetRisksRes{
			ResponseJson: myhttp.ResponseJson{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			},
			Risks: []*Risk{},
		}, err
	}

	return &GetRisksRes{
		ResponseJson: myhttp.ResponseJson{
			StatusCode: http.StatusOK,
			Message:    http.StatusText(http.StatusOK),
		},
		Risks: risks,
	}, nil
}
