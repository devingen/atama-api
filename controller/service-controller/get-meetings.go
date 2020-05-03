package service_controller

import (
	coremodel "github.com/devingen/api-core/model"
	"github.com/devingen/atama-api/dto"
	"github.com/devingen/atama-api/local-data"
)

func (controller ServiceController) GetMeetings(request *dto.GetMeetingsRequest) (*dto.GetMeetingsResponse, error) {

	// TODO remove this when the collection configuration is stored in the database
	collectionConfiguration := local_data.PairCollectionConfiguration.Clone()
	collectionConfiguration.DataModel["space"] = coremodel.DBRef{
		Ref: "organisations",
		ID:  request.SpaceID,
	}

	results, err := controller.Service.GetMeetings(
		request.DatabaseName,
		request.PivotReferenceFieldID,
		request.PivotReferenceFieldValue,
		collectionConfiguration,
	)

	return &dto.GetMeetingsResponse{Results: results}, err
}
