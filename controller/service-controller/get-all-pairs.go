package service_controller

import (
	coremodel "github.com/devingen/api-core/model"
	"github.com/devingen/atama-api/dto"
	local_data "github.com/devingen/atama-api/local-data"
)

func (controller ServiceController) GetAllPairs(request *dto.GetAllPairsRequest) (dto.GetAllPairsResponse, error) {

	// TODO remove this when the collection configuration is stored in the database
	collectionConfiguration := local_data.PairCollectionConfiguration.Clone()
	collectionConfiguration.DataModel["space"] = coremodel.DBRef{
		Ref: "organisations",
		ID:  request.SpaceID,
	}

	results, meta, err := controller.Service.GetAllPairs(
		request.DatabaseName,
		collectionConfiguration,
		local_data.AllPairsViewConfiguration,
		request.Query,
	)

	return dto.GetAllPairsResponse{Results: results, Meta: meta}, err
}
