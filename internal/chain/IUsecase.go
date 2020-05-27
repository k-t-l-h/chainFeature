package chain

import "MarkovChain/internal/models"

type IUsecase interface {
	MakeModel([]models.InputMessage) string
}
