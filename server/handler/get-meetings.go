package handler

import (
	"github.com/devingen/api-core/server"
	"github.com/devingen/api-core/util"
	"github.com/devingen/atama-api/dto"
	"github.com/gorilla/mux"
	"net/http"
)

func (handler ServerHandler) GetMeetings(w http.ResponseWriter, r *http.Request) {

	pathVariables := mux.Vars(r)

	result, err := handler.Controller.GetMeetings(&dto.GetMeetingsRequest{
		DatabaseName:             pathVariables["base"],
		SpaceID:                  pathVariables["space-id"],
		PivotReferenceFieldID:    r.URL.Query().Get("pivot-field"),
		PivotReferenceFieldValue: r.URL.Query().Get("pivot-value"),
	})
	response, err := util.BuildResponse(http.StatusOK, result, err)
	server.ReturnResponse(w, response, err)
}
