package rest_test

//Creates a new testing package to use black box testing        ❶

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"devOps00Boots/handlers/rest"
	//Imports a rest package for testing ❷
)

func TestTranslateAPI(t *testing.T) {

	// Arrange
	rr := httptest.NewRecorder()
	//Creates an HTTP recorder that will be used for assertion ❸
	req, _ := http.NewRequest("GET", "/hello", nil)
	//Creates a new request against a given endpoint with no body content ❹

	handler := http.HandlerFunc(rest.TranslateHandler)
	//Registers a handler to test against ❺

	// Act
	handler.ServeHTTP(rr, req)
	//Serves the content to pass through the handler for a response based on the request ❻

	// Assert
	if rr.Code != http.StatusOK {
		//Checks the status code from the response ❼
		t.Errorf(`expected status 200 but received %d`, rr.Code)
	}

	var resp rest.Resp
	json.Unmarshal(rr.Body.Bytes(), &resp)
	//Decodes the body of the response into a struct to be asserted ❽

	if resp.Language != "english" {
		t.Errorf(`expected language "english" but received %s`,
			resp.Language)
	}

	if resp.Translation != "hello" {
		t.Errorf(`expected Translation "hello" but received %s`,
			resp.Translation)
	}
}
