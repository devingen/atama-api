package dto

import "github.com/devingen/api-core/model"

type GetAllPairsRequest struct {
	DatabaseName string
	SpaceID      string
	Query        *model.BasicQueryConfig
}

type GetAllPairsResponse struct {
	Results []*model.DataModel `json:"results"`
	Meta    *model.Meta        `json:"meta"`
}
