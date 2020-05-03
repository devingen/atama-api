package dto

import "github.com/devingen/api-core/model"

type GetMeetingsRequest struct {
	DatabaseName             string
	SpaceID                  string
	PivotReferenceFieldID    string
	PivotReferenceFieldValue string
}

type GetMeetingsResponse struct {
	Results []*model.DataModel `json:"results"`
}
