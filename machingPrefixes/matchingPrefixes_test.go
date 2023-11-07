package machingPrefixes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetMatchingPrefixRecord(t *testing.T) {
	// Create a new Gin router
	router := gin.New()

	// Define a mock implementation of GetMatchingPrefixesService
	mockGcps := &mockGetMatchingPrefixesService{}
	// Define a mock implementation of HealthCheckService
	mockHcs := &mockHealthCheckService{}

	// Register the endpoint for testing
	RegisterMatchingPrefixesformEndPoint(router, mockGcps, mockHcs)

	// Create a request and recorder
	req, _ := http.NewRequest("GET", "/maching-prefixes/1ci88", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	// Assert the response status code
	assert.Equal(t, http.StatusOK, w.Code)
}

// Mock implementation of GetCustomerPlatformService
type mockGetMatchingPrefixesService struct{}

func (m *mockGetMatchingPrefixesService) GetMatcherPrefixesRecords(c *gin.Context, prefix string) (string, error) {
	// Implement your mock logic here
	return "1ci88", nil
}

// Mock implementation of HealthCheckService
type mockHealthCheckService struct{}

func (m *mockHealthCheckService) HealthCheck(c *gin.Context) {
	// Implement your mock logic here
}
