// Taken from
// https://markjberger.com/testing-web-apps-in-golang/

package utils

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type HandleTester func(
	method string,
	jsonData string,
) *httptest.ResponseRecorder

func GenerateHandleTester(t *testing.T, handleFunc http.HandlerFunc) HandleTester {

	// Given a method type ("GET", "POST", etc) and
	// parameters, serve the response against the handler and
	// return the ResponseRecorder.

	return func(
		method string,
		jsonData string,
	) *httptest.ResponseRecorder {

		req, err := http.NewRequest(
			method,
			"",
			strings.NewReader(jsonData),
		)
		if err != nil {
			t.Errorf("%v", err)
		}

		req.Header.Set("Content-Type", "application/json")
		req.Body.Close()

		w := httptest.NewRecorder()
		handleFunc.ServeHTTP(w, req)
		return w
	}
}
