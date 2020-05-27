package delivery

import (
	"MarkovChain/internal/chain"
	"MarkovChain/internal/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Handler struct {
	Usecase chain.IUsecase
}

func NewHandler(usecase chain.IUsecase) *Handler {
	return &Handler{
		Usecase: usecase,
	}
}

func (h *Handler) GetMessages(w http.ResponseWriter, r *http.Request) {

	content, _ := ioutil.ReadAll(r.Body)
	var data []models.InputMessage
	json.Unmarshal(content, &data)

	generatedMessage := h.Usecase.MakeModel(data)

	response := models.Response{
		Status: http.StatusOK,
		Body:   generatedMessage,
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}
