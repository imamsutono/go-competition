package handler_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-faker/faker/v4"
	"github.com/kompit-recruitment/backend/handler"
	"github.com/stretchr/testify/require"
)

// Sample endpoint.
// (POST /ping)
func TestHandler_PingPost(t *testing.T) {
	reqUrl := "/ping"
	reqMethod := http.MethodPost

	h := handler.New()

	testcases := []struct {
		name               string
		reqBody            interface{}
		expectedStatusCode int
	}{
		{
			name:               "should return 400 if message is not provided",
			reqBody:            map[string]interface{}{},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "should return 400 if message does not start with 'pong'",
			reqBody:            map[string]interface{}{"ping": "hello"},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "should return 200 if message starts with 'pong'",
			reqBody:            map[string]interface{}{"ping": "pong" + faker.Word()},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			httpTest := httptest.NewRecorder()
			ginCtx, _ := gin.CreateTestContext(httpTest)
			// convert request body to json
			jsonReqBody, _ := json.Marshal(tc.reqBody)
			ginCtx.Request = httptest.NewRequest(
				reqMethod, reqUrl, io.NopCloser(bytes.NewBuffer(jsonReqBody)))

			// WHEN
			h.PingPost(ginCtx)

			// THEN
			require.Equal(t, tc.expectedStatusCode, httpTest.Code)
		})
	}
}

// TODO: ASSIGNMENT 6.1
// Creates competition.
// (POST /competitions)
func TestHandler_CompetitionsPost(t *testing.T) {
	t.Skip("TODO: Implement me")
}

// TODO: ASSIGNMENT 6.2
// Returns a competition with the given id.
// (GET /competitions/{id})
func TestHandler_CompetitionsIdGet(t *testing.T) {
	t.Skip("TODO: Implement me")
}

// TODO: ASSIGNMENT 6.3
// Joins a competition with the given id.
// (POST /competitions/{id}/join)
func TestHandler_CompetitionsIdJoinPost(t *testing.T) {
	t.Skip("TODO: Implement me")
}

// TODO: ASSIGNMENT 6.4
// Registers a new user.
// (POST /register)
func TestHandler_RegisterPost(t *testing.T) {
	t.Skip("TODO: Implement me")
}

// TODO: ASSIGNMENT 6.5
// Returns a user with the given username.
// (GET /users/{username})
func TestHandler_UsersUsernameGet(t *testing.T) {
	t.Skip("TODO: Implement me")
}
