package tests

import (
	"net/http"
	"testing"

	"github.com/danmaciel/temperatura_por_cep/internal/entity"
	httpClient "github.com/danmaciel/temperatura_por_cep/internal/infra/http"
	rules "github.com/danmaciel/temperatura_por_cep/internal/rules"
	"github.com/stretchr/testify/assert"
)

func TestCepNotFound(t *testing.T) {
	clientHttp := httpClient.NewHttpClient()

	cep := rules.NewCepRules(clientHttp)

	var c entity.ViaCep
	r := cep.Exec("01153001", &c)

	assert.Equal(t, r.Code, http.StatusNotFound)
	assert.Equal(t, r.Message, "can not find zipcode")
}
