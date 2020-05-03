package integrationtests

import (
	coremodel "github.com/devingen/api-core/model"
	"github.com/devingen/atama-api/model"
)

var viewConfiguration = &model.ViewConfiguration{
	DataModel: coremodel.DataModel{
		"type": model.ViewTypePairs,
		"collectionConfiguration": coremodel.DBRef{
			Ref: "collection-configurations",
			ID:  "1",
		},
		"fields": []coremodel.Field{
			coremodel.NewReference(
				"subscribed",
				"memberships",
				true,
			).SetFields([]coremodel.Field{
				coremodel.NewReference(
					"user",
					"users",
					true,
				).SetFields([]coremodel.Field{
					coremodel.New(coremodel.FieldTypeText, "firstName"),
					coremodel.New(coremodel.FieldTypeText, "lastName"),
				}).ToField(),
			}).ToField(),
			coremodel.NewReference(
				"subscriber",
				"memberships",
				true,
			).SetFields([]coremodel.Field{
				coremodel.NewReference(
					"user",
					"users",
					true,
				).SetFields([]coremodel.Field{
					coremodel.New(coremodel.FieldTypeText, "firstName"),
					coremodel.New(coremodel.FieldTypeText, "lastName"),
				}).ToField(),
			}).ToField(),
		},
	},
}

var pairConfiguration = &model.CollectionConfiguration{
	DataModel: coremodel.DataModel{
		"collection": "subscriptions",
		"space": coremodel.DBRef{
			Ref: "organisations",
			ID:  "1",
		},
		"fields": []coremodel.Field{
			coremodel.New(coremodel.FieldTypeText, "status"),
			coremodel.NewReference(
				"subscribed",
				"memberships",
				true,
			).SetFields([]coremodel.Field{
				coremodel.NewReference(
					"user",
					"users",
					true,
				).ToField(),
			}).ToField(),
			coremodel.NewReference(
				"subscriber",
				"memberships",
				true,
			).SetFields([]coremodel.Field{
				coremodel.NewReference(
					"user",
					"users",
					true,
				).ToField(),
			}).ToField(),
		},
	},
}
