package risks

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RisksController interface {
	CreateRisk(ctx *gin.Context)
	GetRiskById(ctx *gin.Context)
	GetRisks(ctx *gin.Context)
}

type RisksControllerImplementation struct {
	service RisksService
}

func NewRisksController() RisksController {
	risksService := NewRisksService()
	return &RisksControllerImplementation{
		service: risksService,
	}
}

func (controller *RisksControllerImplementation) CreateRisk(ctx *gin.Context) {

	var createRiskReq CreateRiskReq

	err := ctx.ShouldBindJSON(&createRiskReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Something Went Wrong")
	}

	fmt.Println("CreateRiskReq : ", createRiskReq)

	risk, err := controller.service.CreateRisk(ctx, &createRiskReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusCreated, risk)
}

func (controller *RisksControllerImplementation) GetRiskById(ctx *gin.Context) {
	id := ctx.Param("id")
	risk, err := controller.service.GetRiskById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, risk)
}

func (controller *RisksControllerImplementation) GetRisks(ctx *gin.Context) {
	risks := controller.service.GetRisks(ctx)
	ctx.JSON(http.StatusOK, risks)
}
