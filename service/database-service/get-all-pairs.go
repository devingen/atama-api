package database_service

import (
	coremodel "github.com/devingen/api-core/model"
	"github.com/devingen/atama-api/model"
	dataservice "github.com/devingen/data-api/service"
)

func (service DatabaseService) GetAllPairs(
	base string,
	collectionConfiguration *model.CollectionConfiguration,
	viewConfiguration *model.ViewConfiguration,
	basicQuery *coremodel.BasicQueryConfig,
) ([]*coremodel.DataModel, *coremodel.Meta, error) {

	// create a Data API Database AtamaService
	c := dataservice.NewDatabaseService(service.Database)

	return c.Query(
		base,
		collectionConfiguration.GetCollection(),
		&coremodel.QueryConfig{
			Filter: &coremodel.Filter{
				Comparison: coremodel.ComparisonEq,
				FieldId:    "organisation._id",
				FieldValue: collectionConfiguration.GetSpaceID(),
			},
			Limit:  basicQuery.Limit,
			Skip:   basicQuery.Skip,
			Sort:   basicQuery.Sort,
			Fields: viewConfiguration.GetFields(),
		},
	)
}
