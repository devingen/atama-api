package handler

import (
	"github.com/devingen/atama-api/controller"
	"github.com/gorilla/mux"
	"net/http"
)

type ServerHandler struct {
	Controller controller.AtamaController
	Router     *mux.Router
}

func NewHttpServiceHandler(atamaController controller.AtamaController) ServerHandler {
	handler := ServerHandler{Controller: atamaController}

	handler.Router = mux.NewRouter()
	handler.Router.HandleFunc("/{base}/{space-id}/all-pairs", handler.GetAllPairs).Methods(http.MethodGet)
	handler.Router.HandleFunc("/{base}/{space-id}/pairs", handler.GetPairs).Methods(http.MethodGet)
	handler.Router.HandleFunc("/{base}/{space-id}/meeting-structure", handler.GetMeetingStructure).Methods(http.MethodGet)
	handler.Router.HandleFunc("/{base}/{space-id}/meetings", handler.GetMeetings).Methods(http.MethodGet)

	return handler
}
