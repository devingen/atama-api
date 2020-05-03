package handler

import (
	coremodel "github.com/devingen/api-core/model"
	"github.com/devingen/api-core/server"
	"github.com/devingen/api-core/util"
	"github.com/devingen/atama-api/dto"
	"github.com/gorilla/mux"
	"net/http"
)

func (handler ServerHandler) GetPairs(w http.ResponseWriter, r *http.Request) {

	pathVariables := mux.Vars(r)

	result, err := handler.Controller.GetPairs(&dto.GetPairsRequest{
		DatabaseName:             pathVariables["base"],
		SpaceID:                  pathVariables["space-id"],
		PivotReferenceFieldID:    r.URL.Query().Get("pivot-field"),
		PivotReferenceFieldValue: r.URL.Query().Get("pivot-value"),
		Query: &coremodel.BasicQueryConfig{
			Limit: 1000,
			Skip:  0,
			Sort:  nil,
		},
	})
	response, err := util.BuildResponse(http.StatusOK, result, err)
	server.ReturnResponse(w, response, err)
}
