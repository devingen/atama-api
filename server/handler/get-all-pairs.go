package handler

import (
	coremodel "github.com/devingen/api-core/model"
	"github.com/devingen/api-core/server"
	"github.com/devingen/api-core/util"
	"github.com/devingen/atama-api/dto"
	"github.com/gorilla/mux"
	"net/http"
)

func (handler ServerHandler) GetAllPairs(w http.ResponseWriter, r *http.Request) {

	pathVariables := mux.Vars(r)

	result, err := handler.Controller.GetAllPairs(&dto.GetAllPairsRequest{
		DatabaseName: pathVariables["base"],
		SpaceID:      pathVariables["space-id"],
		Query: &coremodel.BasicQueryConfig{
			Limit: 1000,
			Skip:  0,
			Sort:  nil,
		},
	})
	response, err := util.BuildResponse(http.StatusOK, result, err)
	server.ReturnResponse(w, response, err)
}
