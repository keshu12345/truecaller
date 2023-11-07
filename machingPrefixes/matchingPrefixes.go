package machingPrefixes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keshu12345/truecaller/docs"
	"github.com/keshu12345/truecaller/util"
	logger "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	samplePrefixesFile = "sample_prefixes.txt"
)

var gcpsInterafce GetMatcherPrefixesService
var healthCheckInterface HealthCheckService

// Endpoint
func RegisterMatchingPrefixesformEndPoint(g *gin.Engine, gcps GetMatcherPrefixesService, hcs HealthCheckService) {
	docs.SwaggerInfo.BasePath = "/maching-prefixes"
	matchingPrefixRecord := g.Group("/maching-prefixes")
	{
		matchingPrefixRecord.GET("/:prefix_name", GetMatchingPrefixRecord)
		gcpsInterafce = gcps
		healthCheckInterface = hcs

	}
	g.GET("/healthcheck", registerHealthCheck)
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// GetMatchingPrefixRecord godoc
// @Summary      Show an GetMatchingPrefixRecord
// @Description  get string by prefix_name
// @Tags         GetMatchingPrefixRecord
// @Accept       json
// @Produce      json
// @Param        prefix_name   path      int  true  "prefix_name"
// @QueryParams  Key and Value
// @Success      200  {Return raw Object of MatchingPrefixRecord }  gcpsInterafce.SearchLongestPrefix
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /{prefix_name} [get]
func GetMatchingPrefixRecord(c *gin.Context) {

	prefixes, err := util.MatcherPrefixesList(samplePrefixesFile)
	if err != nil {
		logger.Errorln("Unable to parse the sample_prefies.txt file")
	}
	matcherDataStore := &getMatcherPrefixesStrore{}
	matcherDataStore, err = buildMatcherPrefixes(prefixes)
	if err != nil {
		logger.Errorln("Unable to build MatcherPrefixes store data structure")
	}
	gcpsInterafce = matcherDataStore
	prefix_search_input := c.Param("prefix_name")
	prefixMatchingResult, err := gcpsInterafce.GetMatcherPrefixesRecords(c, prefix_search_input)
	if err != nil {
		logger.New().Errorln("in GetMatchingPrefixRecord unable to find logestPrefixes")
	}
	c.JSON(http.StatusOK, prefixMatchingResult)

}

func registerHealthCheck(c *gin.Context) {
	healthCheckInterface.HealthCheck(c)
}
