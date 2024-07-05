package controllers_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/willchat-ofc/api-willchat-golang/internal/presentation/protocols"
)

func verifyHttpResponse(t *testing.T, httpResponse *protocols.HttpResponse, expectedStatusCode int, expectedError string) {
	require.Equal(t, httpResponse.StatusCode, expectedStatusCode)

	var responseBody protocols.ErrorResponse
	err := json.NewDecoder(httpResponse.Body).Decode(&responseBody)
	require.NoError(t, err)
	require.Equal(t, responseBody.Error, expectedError)
}
