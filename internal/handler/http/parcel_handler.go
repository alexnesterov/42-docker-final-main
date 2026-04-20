package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Yandex-Practicum/42-docker-final/internal/domain"
)

type ParcelHandler struct {
	service domain.ParcelService
}

func NewParcelHandler(service domain.ParcelService) *ParcelHandler {
	return &ParcelHandler{service: service}
}

func (h *ParcelHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /parcels", h.register)
	mux.HandleFunc("GET /parcels/{client}", h.getByClient)
	mux.HandleFunc("PUT /parcels/{number}/status", h.nextStatus)
	mux.HandleFunc("PUT /parcels/{number}/address", h.changeAddress)
	mux.HandleFunc("DELETE /parcels/{number}", h.delete)
}

func (h *ParcelHandler) register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Client  int    `json:"client"`
		Address string `json:"address"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	parcel, err := h.service.Register(req.Client, req.Address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, parcel)
}

func (h *ParcelHandler) getByClient(w http.ResponseWriter, r *http.Request) {
	client, err := strconv.Atoi(r.PathValue("client"))
	if err != nil {
		http.Error(w, "invalid client id", http.StatusBadRequest)
		return
	}

	parcels, err := h.service.GetByClient(client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, parcels)
}

func (h *ParcelHandler) nextStatus(w http.ResponseWriter, r *http.Request) {
	number, err := strconv.Atoi(r.PathValue("number"))
	if err != nil {
		http.Error(w, "invalid parcel number", http.StatusBadRequest)
		return
	}

	if err := h.service.NextStatus(number); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ParcelHandler) changeAddress(w http.ResponseWriter, r *http.Request) {
	number, err := strconv.Atoi(r.PathValue("number"))
	if err != nil {
		http.Error(w, "invalid parcel number", http.StatusBadRequest)
		return
	}

	var req struct {
		Address string `json:"address"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.ChangeAddress(number, req.Address); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ParcelHandler) delete(w http.ResponseWriter, r *http.Request) {
	number, err := strconv.Atoi(r.PathValue("number"))
	if err != nil {
		http.Error(w, "invalid parcel number", http.StatusBadRequest)
		return
	}

	if err := h.service.Delete(number); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
