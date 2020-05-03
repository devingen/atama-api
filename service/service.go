package service

import (
	coremodel "github.com/devingen/api-core/model"
	"github.com/devingen/atama-api/model"
)

type AtamaService interface {
	GetAllPairs(base string, collectionConfiguration *model.CollectionConfiguration, viewConfiguration *model.ViewConfiguration, basicQuery *coremodel.BasicQueryConfig) ([]*coremodel.DataModel, *coremodel.Meta, error)
	GetPairs(base string, pivotReferenceFieldID string, pivotReferenceFieldValue string, collectionConfiguration *model.CollectionConfiguration, viewConfiguration *model.ViewConfiguration, basicQuery *coremodel.BasicQueryConfig) ([]*coremodel.DataModel, *coremodel.Meta, error)
	GetMeetingStructure(base string, collectionConfiguration *model.CollectionConfiguration) (*coremodel.DataModel, error)
	GetMeetings(base string,pivotReferenceFieldID string, pivotReferenceFieldValue string, collectionConfiguration *model.CollectionConfiguration) ([]*coremodel.DataModel, error)
}
