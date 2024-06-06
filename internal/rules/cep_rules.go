package usecase

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/danmaciel/temperatura_por_cep/internal/entity"
)

type Cep struct {
	http *http.Client
}

func NewCepRules(h *http.Client) *Cep {
	return &Cep{
		http: h,
	}
}

func (c *Cep) Exec(cep string, vc *entity.ViaCep) error {
	newCep := strings.ReplaceAll(cep, "-", "")

	res, err := c.http.Get(fmt.Sprintf("https://viacep.com.br/ws/%v/json/", newCep))

	if err != nil {
		return err
	}

	json.NewDecoder(res.Body).Decode(&vc)

	return nil
}

func (c *Cep) IsCepValid(cep string) bool {
	newCep := strings.ReplaceAll(cep, "-", "")
	return len(newCep) == 8
}
