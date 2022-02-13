package session

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/chack93/go_base/internal/service/test"
	"github.com/stretchr/testify/assert"
)

func TestGetHello(t *testing.T) {
	ctx, rec := test.Request(
		http.MethodGet,
		"/api/app_name/session/:id",
		[]string{"1"},
		nil,
	)
	if assert.NoError(t, HandlerRead(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var response Session
		expected := Session{}
		assert.Nil(t, json.Unmarshal(rec.Body.Bytes(), &response))
		test.CleanModel(&expected.Model, &response.Model)
		assert.Equal(t, expected, response)
	}
}
