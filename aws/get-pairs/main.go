package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	coreaws "github.com/devingen/api-core/aws"
	coremodel "github.com/devingen/api-core/model"
	"github.com/devingen/api-core/util"
	"github.com/devingen/atama-api/aws"
	"github.com/devingen/atama-api/controller"
	"github.com/devingen/atama-api/dto"
	"github.com/devingen/atama-api/service"
	"net/http"
)

func main() {

	db := aws.GetDatabase()
	databaseService := service.NewDatabaseService(db)
	serviceController := controller.NewServiceController(databaseService)

	lambda.Start(func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

		result, err := serviceController.GetPairs(&dto.GetPairsRequest{
			DatabaseName:             req.PathParameters["base"],
			SpaceID:                  req.PathParameters["space-id"],
			PivotReferenceFieldID:    req.QueryStringParameters["pivot-field"],
			PivotReferenceFieldValue: req.QueryStringParameters["pivot-value"],
			Query: &coremodel.BasicQueryConfig{
				Limit: 1000,
				Skip:  0,
				Sort:  nil,
			},
		})
		response, err := util.BuildResponse(http.StatusOK, result, err)
		return coreaws.AdaptResponse(response, err)
	})
}
