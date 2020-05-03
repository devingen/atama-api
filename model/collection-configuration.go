package model

import coreModel "github.com/devingen/api-core/model"

type CollectionType string

const (
	CollectionTypePair CollectionType = "pair"
)

type CollectionConfiguration struct {
	coreModel.DataModel
}

func (c CollectionConfiguration) GetCollection() string {
	return c.GetString("collection")
}

func (c CollectionConfiguration) GetSpaceID() string {
	return c.GetDBRef("space").ID
}

func CopyMap(m map[string]interface{}) map[string]interface{} {
	cp := make(map[string]interface{})
	for k, v := range m {
		vm, ok := v.(map[string]interface{})
		if ok {
			cp[k] = CopyMap(vm)
		} else {
			cp[k] = v
		}
	}

	return cp
}

func (c CollectionConfiguration) Clone() *CollectionConfiguration {
	return &CollectionConfiguration{DataModel: CopyMap(c.DataModel)}
}
