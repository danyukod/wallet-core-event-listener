package web

import (
	"encoding/json"
	"fmt"
	"github.com/danyukod/wallet-core-event-listener/internal/usecase/get_balance"
	"github.com/go-chi/chi"
	"net/http"
)

type WebBalanceHandler struct {
	GetBalanceUseCase get_balance.GetBalanceUseCase
}

func NewWebBalanceHandler(getBalanceUseCase get_balance.GetBalanceUseCase) *WebBalanceHandler {
	return &WebBalanceHandler{
		GetBalanceUseCase: getBalanceUseCase,
	}
}

func (h *WebBalanceHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	var dto get_balance.GetBalanceInputDTO
	accountID := chi.URLParam(r, "account_id")
	if accountID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return

	}
	dto.AccountID = accountID

	output, err := h.GetBalanceUseCase.Execute(dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}
