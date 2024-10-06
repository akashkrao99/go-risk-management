// risks/tests/service_test.go
package risks

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/akashkrao99/go-sample-http/internal/risks"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

// MockRisksRepository is a mock implementation of the RisksRepository interface.
type MockRisksRepository struct {
    mock.Mock
}

func (m *MockRisksRepository) CreateRisk(ctx *gin.Context, risk *risks.Risk) (*risks.Risk, error) {
    args := m.Called(ctx, risk)
    return args.Get(0).(*risks.Risk), args.Error(1)
}

func (m *MockRisksRepository) GetRiskById(ctx *gin.Context, id string) (*risks.Risk, error) {
    args := m.Called(ctx, id)
    return args.Get(0).(*risks.Risk), args.Error(1)
}

func (m *MockRisksRepository) GetRisks(ctx *gin.Context) ([]*risks.Risk, error) {
    args := m.Called(ctx)
    return args.Get(0).([]*risks.Risk), args.Error(1)
}

func TestCreateRisk(t *testing.T) {
    mockRepo := new(MockRisksRepository)
    service := risks.NewRisksServiceImplementation(mockRepo)
    ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

    req := &risks.CreateRiskReq{
       Title:       "Test Risk",
       Description: "This is description.",
       Status:      risks.STATUS_OPEN,
    }

    // Test valid creation
    mockRepo.On("CreateRisk", ctx, mock.AnythingOfType("*risks.Risk")).Return(&risks.Risk{Id: uuid.New().String()}, nil)

    resp, err := service.CreateRisk(ctx, req)
    assert.NoError(t, err)
    assert.Equal(t, http.StatusCreated, resp.ResponseJson.StatusCode)

    // Test error on validation
    reqInvalid := &risks.CreateRiskReq{
       Title:       "Shrt",
       Description: "Desc",
       Status:      risks.STATUS_OPEN,
    }

    respInvalid, err := service.CreateRisk(ctx, reqInvalid)
    assert.Error(t, err)
    assert.Equal(t, http.StatusBadRequest, respInvalid.ResponseJson.StatusCode)
}

func TestGetRiskById(t *testing.T) {
    mockRepo := new(MockRisksRepository)
    service := risks.NewRisksServiceImplementation(mockRepo)
    ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

    riskID := uuid.New().String()

    // Test valid retrieval
    mockRepo.On("GetRiskById", ctx, riskID).Return(&risks.Risk{Id: riskID}, nil)

    resp, err := service.GetRiskById(ctx, riskID)
    assert.NoError(t, err)
    assert.Equal(t, http.StatusOK, resp.ResponseJson.StatusCode)

}

func TestGetRisks(t *testing.T) {
    mockRepo := new(MockRisksRepository)
    service := risks.NewRisksServiceImplementation(mockRepo)

    ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

    // Test valid retrieval of risks
    mockRepo.On("GetRisks", ctx).Return([]*risks.Risk{{Id: uuid.New().String()}}, nil)

    resp, err := service.GetRisks(ctx)
    assert.NoError(t, err)
    assert.Equal(t, http.StatusOK, resp.ResponseJson.StatusCode)
    assert.NotEmpty(t, resp.Risks)

}