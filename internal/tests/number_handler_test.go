package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"log/slog"
	"io"

	"github.com/gin-gonic/gin"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"goforge/internal/domain/models"
	"goforge/internal/http-server/handler"
	"goforge/internal/transport/http/request"

)

type MockNumbersProvider struct {
	mock.Mock
}

func (m *MockNumbersProvider) SaveNumber(num models.Number) ([]int, error) {
	args := m.Called(num)
	return args.Get(0).([]int), args.Error(1)
}	
func TestNumberAddSuccess(t *testing.T) {
	
	gin.SetMode(gin.TestMode)
	
	mockProvider := new(MockNumbersProvider)
	h := &handler.Numbers{
		
		Log:   slogDiscard(),
		NumbersProvider: mockProvider,
	}

	router := gin.New()
	router.POST("/add", h.NumberAdd)

	mockProvider.
		On("SaveNumber", models.Number{Value: 3}).
		Return([]int{3}, nil)

	body, _ := json.Marshal(request.Number{Value: 3})
	req := httptest.NewRequest(http.MethodPost, "/add", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `{"numbers":[3]}`, rr.Body.String())
	mockProvider.AssertExpectations(t)
}

func slogDiscard() *slog.Logger {
	return slog.New(slog.NewTextHandler(io.Discard, nil))
}
