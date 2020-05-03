package database_service

import (
	coremodel "github.com/devingen/api-core/model"
	"github.com/devingen/atama-api/model"
	dataservice "github.com/devingen/data-api/service"
)

func (service DatabaseService) GetPairs(
	base string,
	pivotReferenceFieldID string,
	pivotReferenceFieldValue string,
	collectionConfiguration *model.CollectionConfiguration,
	viewConfiguration *model.ViewConfiguration,
	basicQuery *coremodel.BasicQueryConfig,
) ([]*coremodel.DataModel, *coremodel.Meta, error) {

	// create a Data API Database AtamaService
	c := dataservice.NewDatabaseService(service.Database)

	// trim the fields of the pivot reference field
	queryFields := make([]coremodel.Field, 0)
	for _, field := range viewConfiguration.GetFields() {
		if field.GetType() != coremodel.FieldTypeReference {
			queryFields = append(queryFields, field)
			continue
		}

		if field.GetID() == pivotReferenceFieldID {
			fieldWithoutInnerFields := coremodel.Field{}
			for k, v := range field {
				if k != "fields" {
					fieldWithoutInnerFields[k] = v
				}
			}
			queryFields = append(queryFields, fieldWithoutInnerFields)
			continue
		}
		queryFields = append(queryFields, field)
	}

	return c.QueryCollection(
		base,
		collectionConfiguration.GetCollection(),
		&coremodel.QueryConfig{
			Filter: &coremodel.Filter{
				Operator: coremodel.OperatorAnd,
				Filters: []coremodel.Filter{
					{
						Comparison: coremodel.ComparisonEq,
						FieldId:    "organisation._id",
						FieldValue: collectionConfiguration.GetSpaceID(),
					},
					{
						Comparison: coremodel.ComparisonEq,
						FieldId:    pivotReferenceFieldID + "._id",
						FieldValue: pivotReferenceFieldValue,
					},
					{
						Comparison: coremodel.ComparisonEq,
						FieldId:    "status",
						FieldValue: "active",
					},
				},
			},
			Limit:  basicQuery.Limit,
			Skip:   basicQuery.Skip,
			Sort:   basicQuery.Sort,
			Fields: queryFields,
		},
	)
}
