package controller

import (
	"github.com/devingen/atama-api/dto"
)

type AtamaController interface {
	GetAllPairs(request *dto.GetAllPairsRequest) (dto.GetAllPairsResponse, error)
	GetPairs(request *dto.GetPairsRequest) (dto.GetPairsResponse, error)
	GetMeetingStructure(request *dto.GetMeetingStructureRequest) (*dto.GetMeetingStructureResponse, error)
	GetMeetings(request *dto.GetMeetingsRequest) (*dto.GetMeetingsResponse, error)
}
