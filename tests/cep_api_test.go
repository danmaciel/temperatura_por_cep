package tests

import (
	"testing"

	"github.com/danmaciel/temperatura_por_cep/internal/entity"
	httpClient "github.com/danmaciel/temperatura_por_cep/internal/infra/http"
	rules "github.com/danmaciel/temperatura_por_cep/internal/rules"
	"github.com/stretchr/testify/assert"
)

func TestCepApi(t *testing.T) {

	clientHttp := httpClient.NewHttpClient()

	cep := rules.NewCepRules(clientHttp)

	var cepModel entity.ViaCep
	err := cep.Exec("86010-180", &cepModel)

	assert.Nil(t, err)
	assert.NotNil(t, cepModel)
	assert.Equal(t, cepModel.City, "Londrina")
}
