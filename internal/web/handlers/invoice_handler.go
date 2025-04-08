package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/marcofilho/go-api-payment-gateway/internal/domain"
	"github.com/marcofilho/go-api-payment-gateway/internal/dto"
	"github.com/marcofilho/go-api-payment-gateway/internal/service"
)

type InvoiceHandler struct {
	invoiceService *service.InvoiceService
	accountService *service.AccountService
}

func NewInvoiceHandler(invoiceService *service.InvoiceService, accountService *service.AccountService) *InvoiceHandler {
	return &InvoiceHandler{
		invoiceService: invoiceService,
		accountService: accountService,
	}
}

func (h *InvoiceHandler) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateInvoiceDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	input.APIKey = r.Header.Get("X-API-KEY")

	output, err := h.invoiceService.CreateInvoice(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (h *InvoiceHandler) GetInvoiceByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	apiKey := r.Header.Get("X-API-KEY")

	output, err := h.invoiceService.GetById(id, apiKey)
	if err != nil {
		switch err {
		case domain.ErrorAccountNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		case domain.ErrorInvoiceNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		case domain.ErrorUnauthorizedAccess:
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (h *InvoiceHandler) GetInvoicesByAccountID(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-API-KEY")
	if apiKey == "" {
		http.Error(w, "API key is required", http.StatusUnauthorized)
		return
	}

	invoicesOutput, err := h.invoiceService.GetByAccountAPIKey(apiKey)
	if err != nil {
		switch err {
		case domain.ErrorAccountNotFound:
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(invoicesOutput)
}
