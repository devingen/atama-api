package database_service

import (
	coremodel "github.com/devingen/api-core/model"
	"github.com/devingen/atama-api/model"
	dataservice "github.com/devingen/data-api/service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (service DatabaseService) GetMeetings(
	base string,
	pivotReferenceFieldID string,
	pivotReferenceFieldValue string,
	collectionConfiguration *model.CollectionConfiguration,
) ([]*coremodel.DataModel, error) {

	// create a Data API Database AtamaService
	c := dataservice.NewDatabaseService(service.Database)

	// trim the fields of all references to get only IDs
	queryFields := make([]coremodel.Field, 0)
	for _, field := range collectionConfiguration.GetFields() {
		if field.GetType() != coremodel.FieldTypeReference {
			continue
		}

		queryFields = append(queryFields, coremodel.Field{
			"id": field.GetID(),
		})
	}

	pairs, _, err := c.Query(
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
			Limit:  1000,
			Skip:   0,
			Fields: queryFields,
		},
	)
	if err != nil {
		return nil, err
	}

	otherGuestsCondition := make(bson.A, 0)
	for _, pair := range pairs {
		guestId := ""
		if pivotReferenceFieldID == "subscriber" {
			guestId = pair.GetChildDataModel("subscribed").GetString("_id")
		} else {
			guestId = pair.GetChildDataModel("subscriber").GetString("_id")
		}
		otherGuestsCondition = append(otherGuestsCondition, bson.M{"guests": bson.M{"$elemMatch": bson.M{"_id": guestId}}})
	}

	events := make([]*coremodel.DataModel, 0)
	err = c.Database.Query(base, "events", bson.M{
		"$and": bson.A{
			bson.M{"status": "active"},
			bson.M{"guests": bson.M{"$elemMatch": bson.M{"_id": pivotReferenceFieldValue}}},
			bson.M{"$or": otherGuestsCondition},
		},
	}, func(cur *mongo.Cursor) error {
		var event coremodel.DataModel
		err := cur.Decode(&event)
		if err != nil {
			return err
		}
		events = append(events, &coremodel.DataModel{
			"guests":                event["guests"],
			"meetingGoalIdentifier": event["meetingGoalIdentifier"],
			"startDate":             event["startDate"],
			"endDate":               event["endDate"],
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	return events, nil
}
