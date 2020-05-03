package local_data

import (
	coremodel "github.com/devingen/api-core/model"
	"github.com/devingen/atama-api/model"
)

var AllPairsViewConfiguration = &model.ViewConfiguration{
	DataModel: coremodel.DataModel{
		"type": model.ViewTypePairs,
		"collectionConfiguration": coremodel.DBRef{
			Ref: "collection-configurations",
			ID:  "1",
		},
		"fields": []coremodel.Field{
			coremodel.NewReference(
				"subscribed",
				"users",
				true,
			).SetFields([]coremodel.Field{
				coremodel.New(coremodel.FieldTypeText, "firstName"),
				coremodel.New(coremodel.FieldTypeText, "lastName"),
				coremodel.New(coremodel.FieldTypeText, "imageUrl"),
			}).ToField(),
			coremodel.NewReference(
				"subscriber",
				"users",
				true,
			).SetFields([]coremodel.Field{
				coremodel.New(coremodel.FieldTypeText, "firstName"),
				coremodel.New(coremodel.FieldTypeText, "lastName"),
				coremodel.New(coremodel.FieldTypeText, "imageUrl"),
			}).ToField(),
			coremodel.New(coremodel.FieldTypeText, "status"),
		},
	},
}

var PairsViewConfiguration = &model.ViewConfiguration{
	DataModel: coremodel.DataModel{
		"type": model.ViewTypePairs,
		"collectionConfiguration": coremodel.DBRef{
			Ref: "collection-configurations",
			ID:  "1",
		},
		"fields": []coremodel.Field{
			coremodel.NewReference(
				"subscribed",
				"users",
				true,
			).SetFields([]coremodel.Field{
				coremodel.New(coremodel.FieldTypeText, "firstName"),
				coremodel.New(coremodel.FieldTypeText, "lastName"),
				coremodel.New(coremodel.FieldTypeText, "imageUrl"),
			}).ToField(),
			coremodel.NewReference(
				"subscriber",
				"users",
				true,
			).SetFields([]coremodel.Field{
				coremodel.New(coremodel.FieldTypeText, "firstName"),
				coremodel.New(coremodel.FieldTypeText, "lastName"),
				coremodel.New(coremodel.FieldTypeText, "imageUrl"),
			}).ToField(),
		},
	},
}

// This configuration is used for all spaces. Each pair collection should
// have its own configuration set in database. The 'space' field below is
// filled dynamically when this configuration is used.
var PairCollectionConfiguration = &model.CollectionConfiguration{
	DataModel: coremodel.DataModel{
		"type":       model.CollectionTypePair,
		"collection": "subscriptions",
		"space":      coremodel.DBRef{}, // filled by handler
		"fields": []coremodel.Field{
			coremodel.NewReference(
				"organisation",
				"organisations",
				true,
			).ToField(),
			coremodel.NewReference(
				"createdBy",
				"users",
				true,
			).ToField(),
			coremodel.NewReference(
				"subscribed",
				"users",
				true,
			).ToField(),
			coremodel.NewReference(
				"subscriber",
				"users",
				true,
			).ToField(),
			coremodel.New(coremodel.FieldTypeText, "status"),
			coremodel.New(coremodel.FieldTypeBoolean, "isFollowing"),
		},
	},
}
