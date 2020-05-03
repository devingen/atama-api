package model

import coreModel "github.com/devingen/api-core/model"

type ViewType string

const (
	ViewTypePairs ViewType = "pairs"
)

type ViewConfiguration struct {
	coreModel.DataModel
}
