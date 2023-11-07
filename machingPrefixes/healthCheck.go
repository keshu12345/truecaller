package machingPrefixes

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keshu12345/truecaller/util"
	logger "github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var errMatchingPrefixesConnectionFailed = errors.New("ERR_PREFIXIES_MATCHES_BUILD_FAILED")

// Health Check
type HealthCheckService interface {
	HealthCheck(c *gin.Context)
}

// Inject an object into healthCheckService
type healthCheckService struct {
	fx.In
}

func NewHealthCheckService(heathCheck healthCheckService) HealthCheckService {
	return healthCheckService{}
}

func (healthCheck healthCheckService) HealthCheck(ctx *gin.Context) {
	prefixes, err := util.MatcherPrefixesList(samplePrefixesFile)
	if err != nil {
		logger.Errorln("Unable to parse the sample_prefies.txt file")
	}
	_, err = buildMatcherPrefixes(prefixes)
	if err != nil {
		logger.Errorln("Unable to build MatcherPrefixes store data structure")
	}
	if err != nil {
		returnError(ctx, errMatchingPrefixesConnectionFailed, err.Error(), http.StatusInternalServerError)
	}
	returnSuccess(ctx, nil, http.StatusOK, "healthCheck")
}
