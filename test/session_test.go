package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chack93/go_base/internal/domain/session"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestSessionCLRUD(t *testing.T) {
	var createResponse session.Session
	var ctx echo.Context
	var rec *httptest.ResponseRecorder

	// CREATE
	ctx, rec = Request("POST", "/api/go_base/session/", []string{}, nil)
	if assert.NoError(t, session.HandlerCreate(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Nil(t, json.Unmarshal(rec.Body.Bytes(), &createResponse))
	}
	// LIST
	ctx, rec = Request("GET", "/api/go_base/session/", []string{}, nil)
	if assert.NoError(t, session.HandlerList(ctx)) {
		var expected = []session.Session{{
			JoinCode: createResponse.JoinCode,
		}}
		var response []session.Session
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Nil(t, json.Unmarshal(rec.Body.Bytes(), &response))
		assert.Equal(t, len(expected), len(response))
		assert.Equal(t, expected[0].JoinCode, response[0].JoinCode)
	}
	// READ
	ctx, rec = Request("GET", "/api/go_base/session/:id", []string{createResponse.ID.String()}, nil)
	if assert.NoError(t, session.HandlerRead(ctx)) {
		var expected = session.Session{
			JoinCode: createResponse.JoinCode,
		}
		expected.ID = createResponse.ID
		assert.Equal(t, http.StatusOK, rec.Code)
		var response session.Session
		assert.Nil(t, json.Unmarshal(rec.Body.Bytes(), &response))
		CleanModelTS(&expected.Model, &response.Model)
		assert.Equal(t, expected, response)
	}
	// UPDATE
	var update = session.Session{
		JoinCode: uuid.NewString(),
	}
	update.ID = createResponse.ID
	ctx, rec = Request("PUT", "/api/go_base/session/", []string{}, &update)
	if assert.NoError(t, session.HandlerUpdate(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var response session.Session
		assert.Nil(t, json.Unmarshal(rec.Body.Bytes(), &response))
		CleanModelTS(&update.Model, &response.Model)
		assert.Equal(t, update, response)
	}
	// DELETE
	ctx, rec = Request("DELETE", "/api/go_base/session/:id", []string{createResponse.ID.String()}, nil)
	if assert.NoError(t, session.HandlerDelete(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var response session.Session
		assert.Nil(t, json.Unmarshal(rec.Body.Bytes(), &response))
		CleanModelTS(&update.Model, &response.Model)
		assert.Equal(t, update, response)

		ctx, rec = Request("GET", "/api/go_base/session/:id", []string{createResponse.ID.String()}, nil)
		if assert.Error(t, session.HandlerRead(ctx)) {

		}
	}
}
