package tests

import (
	"testing"

	httpClient "github.com/danmaciel/temperatura_por_cep/internal/infra/http"
	rules "github.com/danmaciel/temperatura_por_cep/internal/rules"
	"github.com/stretchr/testify/assert"
)

func TestCepLenght(t *testing.T) {
	clientHttp := httpClient.NewHttpClient()

	cep := rules.NewCepRules(clientHttp)

	r := cep.IsCepValid("1234567")

	assert.False(t, r)

	r = cep.IsCepValid("12345678")

	assert.True(t, r)
}
