package dto

import "github.com/devingen/api-core/model"

type GetPairsRequest struct {
	DatabaseName             string
	SpaceID                  string
	PivotReferenceFieldID    string
	PivotReferenceFieldValue string
	Query                    *model.BasicQueryConfig
}

type GetPairsResponse struct {
	Results []*model.DataModel `json:"results"`
	Meta    *model.Meta        `json:"meta"`
}
