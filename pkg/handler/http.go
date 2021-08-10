package handler

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	plansservice "github.com/karakanb/plans-service-grpc-poc/proto"
	"net/http"
)

type HttpServer struct {
	handler *handler
}

func NewHttpServer(handler *handler) *HttpServer {
	return &HttpServer{handler: handler}
}

func (s *HttpServer) CreatePlan(w http.ResponseWriter, r *http.Request) {
	var planReq plansservice.PlanRequest
	err := json.NewDecoder(r.Body).Decode(&planReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	plan, err := s.handler.CreatePlan(context.Background(), &planReq)
	renderJSON(w, http.StatusCreated, plan)
}

func (s *HttpServer) GetPlan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	getPlanRequest := plansservice.GetPlanRequest{ID: vars["id"]}

	plan, err := s.handler.GetPlan(context.Background(), &getPlanRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	renderJSON(w, http.StatusOK, plan)
}

func renderJSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	if v == nil || code == http.StatusNoContent {
		return
	}

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(true)

	if err := enc.Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
