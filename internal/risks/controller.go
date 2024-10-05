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

	res, _ := controller.service.CreateRisk(ctx, &createRiskReq)
	ctx.JSON(res.ResponseJson.StatusCode, res)
}

func (controller *RisksControllerImplementation) GetRiskById(ctx *gin.Context) {
	id := ctx.Param("id")
	res,_ := controller.service.GetRiskById(ctx, id)
	ctx.JSON(res.ResponseJson.StatusCode, res)
}

func (controller *RisksControllerImplementation) GetRisks(ctx *gin.Context) {
	res,_ := controller.service.GetRisks(ctx)
	ctx.JSON(res.ResponseJson.StatusCode, res)
}
