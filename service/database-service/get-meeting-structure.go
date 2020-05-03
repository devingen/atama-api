package database_service

import (
	coremodel "github.com/devingen/api-core/model"
	"github.com/devingen/atama-api/model"
	dataservice "github.com/devingen/data-api/service"
)

func (service DatabaseService) GetMeetingStructure(base string, collectionConfiguration *model.CollectionConfiguration) (*coremodel.DataModel, error) {

	// create a Data API Database AtamaService
	c := dataservice.NewDatabaseService(service.Database)

	results, _, err := c.Query(
		base,
		"meetingGoalConfigs",
		&coremodel.QueryConfig{
			Filter: &coremodel.Filter{
				Comparison: coremodel.ComparisonEq,
				FieldId:    "organisation._id",
				FieldValue: collectionConfiguration.GetSpaceID(),
			},
			Limit: 1,
			Skip:  0,
			Fields: []coremodel.Field{
				coremodel.New(coremodel.FieldTypeAny, "meetingGoals"),
			},
		},
	)

	if err != nil {
		return nil, err
	} else if len(results) == 0 {
		return nil, nil
	}
	return results[0], nil
}
