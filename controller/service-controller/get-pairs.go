package service_controller

import (
	coremodel "github.com/devingen/api-core/model"
	"github.com/devingen/atama-api/dto"
	"github.com/devingen/atama-api/local-data"
)

func (controller ServiceController) GetPairs(request *dto.GetPairsRequest) (dto.GetPairsResponse, error) {

	// TODO remove this when the collection configuration is stored in the database
	collectionConfiguration := local_data.PairCollectionConfiguration.Clone()
	collectionConfiguration.DataModel["space"] = coremodel.DBRef{
		Ref: "organisations",
		ID:  request.SpaceID,
	}

	results, meta, err := controller.Service.GetPairs(
		request.DatabaseName,
		request.PivotReferenceFieldID,
		request.PivotReferenceFieldValue,
		collectionConfiguration,
		local_data.PairsViewConfiguration,
		request.Query,
	)

	return dto.GetPairsResponse{Results: results, Meta: meta}, err
}
