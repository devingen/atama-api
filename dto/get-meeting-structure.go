package dto

import "github.com/devingen/api-core/model"

type GetMeetingStructureRequest struct {
	DatabaseName string
	SpaceID      string
}

type GetMeetingStructureResponse = *model.DataModel
