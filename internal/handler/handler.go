package handler

import (
	"encoding/json"
	"github.com/rabboni171/transactions-demo/internal/service"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	svc    service.AccountService
	logger *logrus.Logger
}

func NewHandler(svc service.AccountService, logger *logrus.Logger) *Handler {
	return &Handler{
		svc:    svc,
		logger: logger,
	}
}

type TransferRequest struct {
	FromID int `json:"from_id" binding:"required"`
	ToID   int `json:"to_id" binding:"required"`
	Amount int `json:"amount" binding:"required"`
}

func (h *Handler) Transfer(w http.ResponseWriter, r *http.Request) {
	var req TransferRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		h.logger.Error("error while decoding request")
		h.respondWithError(w, http.StatusBadRequest, "invalid json format")
		return
	}

	err := h.svc.Transfer(req.FromID, req.ToID, req.Amount)
	if err != nil {
		h.logger.Info("error while transferring", err)
		h.respondWithError(w, http.StatusInternalServerError, "error transfering funds")
		return
	}

	h.respondWithJSON(w, http.StatusOK, map[string]string{"message": "transfer completed"})
}
