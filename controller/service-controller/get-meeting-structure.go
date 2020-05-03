package service_controller

import (
	coremodel "github.com/devingen/api-core/model"
	"github.com/devingen/atama-api/dto"
	"github.com/devingen/atama-api/model"
)

func (controller ServiceController) GetMeetingStructure(request *dto.GetMeetingStructureRequest) (*dto.GetMeetingStructureResponse, error) {

	meetingStructure, err := controller.Service.GetMeetingStructure(
		request.DatabaseName,
		&model.CollectionConfiguration{
			DataModel: coremodel.DataModel{
				"space": coremodel.DBRef{
					Ref: "organisations",
					ID:  request.SpaceID,
				},
			},
		},
	)

	return &meetingStructure, err
}
