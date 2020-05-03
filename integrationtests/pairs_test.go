package integrationtests

import (
	"github.com/devingen/api-core/database"
	coremodel "github.com/devingen/api-core/model"
	"github.com/devingen/api-core/util"
	"github.com/devingen/atama-api/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
)

type PairsTestSuite struct {
	suite.Suite
	service service.AtamaService
	base    string
}

func TestPairs(t *testing.T) {

	db, err := database.NewDatabaseWithURI("mongodb://localhost")
	if err != nil {
		log.Fatalf("Database connection failed %s", err.Error())
	}

	testSuite := &PairsTestSuite{
		service: service.NewDatabaseService(db),
		base:    "dvn-atama-api-integration-test",
	}

	util.InsertDataFromFile(db, testSuite.base, "organisations")
	util.InsertDataFromFile(db, testSuite.base, "users")
	util.InsertDataFromFile(db, testSuite.base, "memberships")
	util.InsertDataFromFile(db, testSuite.base, "subscriptions")

	suite.Run(t, testSuite)
}

func (suite *PairsTestSuite) TestGetAllPairs() {

	results, _, err := suite.service.GetAllPairs(
		suite.base,
		pairConfiguration,
		viewConfiguration,
		&coremodel.BasicQueryConfig{
			Limit: 10,
			Skip:  0,
			Sort:  nil,
		},
	)
	util.SaveResultFile("get-pairs-in-space", results)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 2, len(results))
}

func (suite *PairsTestSuite) TestGetPairsAsSubscribed() {

	results, _, err := suite.service.GetPairs(
		suite.base,
		"subscribed",
		"1",
		pairConfiguration,
		viewConfiguration,
		&coremodel.BasicQueryConfig{
			Limit: 10,
			Skip:  0,
			Sort:  nil,
		},
	)
	util.SaveResultFile("get-pairs-as-subscribed", results)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 1, len(results))
}

func (suite *PairsTestSuite) TestGetPairsAsSubscriber() {

	results, _, err := suite.service.GetPairs(
		suite.base,
		"subscriber",
		"2",
		pairConfiguration,
		viewConfiguration,
		&coremodel.BasicQueryConfig{
			Limit: 10,
			Skip:  0,
			Sort:  nil,
		},
	)
	util.SaveResultFile("get-pairs-as-subscriber", results)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 1, len(results))
}
