package risks

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

var inMemoryRisksMap map[string]Risk

type RisksRepository interface {
	CreateRisk(ctx *gin.Context, risk *Risk) (*Risk, error)
	GetRiskById(ctx *gin.Context, id string) (*Risk, error)
	GetRisks(ctx *gin.Context) []*Risk
}

type RisksRepositoryImplementation struct{}

func NewRisksRepository() RisksRepository {
	return &RisksRepositoryImplementation{}
}

func (repo *RisksRepositoryImplementation) CreateRisk(ctx *gin.Context, risk *Risk) (*Risk, error) {

	_, ok := inMemoryRisksMap[risk.Id]
	if ok {
		errStr := fmt.Sprintf("Risk Already Exisits for ID : %v", risk.Id)
		fmt.Errorf(errStr)
		return nil, errors.New(errStr)
	}

	// basic data sanitization can be performed here
	risk.sanitize()

	err := risk.isValid()
	if err != nil {
		fmt.Errorf("model validation failed : ", err.Error())
		return nil, err
	}

	inMemoryRisksMap[risk.Id] = *risk

	return risk, nil

}

func (repo *RisksRepositoryImplementation) GetRiskById(ctx *gin.Context, id string) (*Risk, error) {

	value, ok := inMemoryRisksMap[id]
	if !ok {
		errStr := fmt.Sprintf("Risk Does Not Exisits for ID : %v", id)
		fmt.Errorf(errStr)
		return nil, errors.New(errStr)
	}

	return &value, nil

}

func (repo *RisksRepositoryImplementation) GetRisks(ctx *gin.Context) []*Risk {
	// Create a slice to hold the Person structs
	var risks []*Risk

	// Iterate over the map and append the values (Person structs) to the slice
	for _, risk := range inMemoryRisksMap {
		risks = append(risks, &risk)
	}

	return risks
}
