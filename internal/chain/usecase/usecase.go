package usecase

import (
	"MarkovChain/internal/models"
	"github.com/mb-14/gomarkov"
	"regexp"
	"strings"
)

type Usecase struct {
}

func NewUsecase() *Usecase {
	return &Usecase{}
}

func (us *Usecase) MakeModel(data []models.InputMessage) string {

	chain := gomarkov.NewChain(1)

	for _, text := range data {

		sentence := regexp.MustCompile("[(),.?;!]").Split(text.Message, -1)

		for _, st := range sentence {
			strings.Trim(st, st)
			st = strings.ToLower(st)
			words := strings.Split(st, " ")
			chain.Add(words)
		}
	}

	return GenerateText(chain)
}

func GenerateText(chain *gomarkov.Chain) string {
	tokens := []string{gomarkov.StartToken}

	for tokens[len(tokens)-1] != gomarkov.EndToken {
		next, _ := chain.Generate(tokens[(len(tokens) - 1):])
		tokens = append(tokens, next)
	}
	return strings.Join(tokens[1:len(tokens)-1], " ")
}
