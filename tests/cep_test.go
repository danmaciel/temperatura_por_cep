package tests

import (
	"net/http"
	"testing"

	"github.com/danmaciel/temperatura_por_cep/internal/entity"
	httpClient "github.com/danmaciel/temperatura_por_cep/internal/infra/http"
	rules "github.com/danmaciel/temperatura_por_cep/internal/rules"
	"github.com/stretchr/testify/assert"
)

func TestCepValid(t *testing.T) {
	clientHttp := httpClient.NewHttpClient()

	cep := rules.NewCepRules(clientHttp)

	var c entity.ViaCep
	r := cep.Exec("12345678", &c)

	assert.NotEqual(t, r.Code, http.StatusUnprocessableEntity)
	assert.NotEqual(t, r.Message, "invalid zipcode")

	r = cep.Exec("1234567", &c)

	assert.Equal(t, r.Code, http.StatusUnprocessableEntity)
	assert.Equal(t, r.Message, "invalid zipcode")
}
