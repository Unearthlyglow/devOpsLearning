// You will also notice the use of the black box testing approach. This refers to code packages in which tests cannot see inside the code to see how it works. This allows us to write tests that assert behavior and not implementation. Remember that the system under test should be tested on its inputs and outputs and not how it works internally. This also requires you to think of an appropriate interface, or exposed definition for your application and code. The unit you are developing is an abstraction for others to use. Writing good tests helps drive a good interface. Having a good interface is important because once an interface is exposed, you will need to support it in the future, and it will become hard to change.

package translation_test

import (
	// Uses a separate package to provide black box testing
	"devOps00Boots/handlers/rest"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTranslateAPI(t *testing.T) {
	tt := []struct {
		//❶
		Endpoint            string
		StatusCode          int
		ExpectedLanguage    string
		ExpectedTranslation string
	}{
		{
			Endpoint:            "/hello",
			StatusCode:          http.StatusOK,
			ExpectedLanguage:    "english",
			ExpectedTranslation: "hello",
		},
		{
			Endpoint:            "/hello?language=german",
			StatusCode:          http.StatusOK,
			ExpectedLanguage:    "german",
			ExpectedTranslation: "hallo",
		},
		{
			Endpoint:            "/hello?language=dutch",
			StatusCode:          http.StatusNotFound,
			ExpectedLanguage:    "",
			ExpectedTranslation: "",
		},
	}

	handler := http.HandlerFunc(rest.TranslateHandler)
	//❷

	for _, test := range tt {
		// ❸
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", test.Endpoint, nil)

		handler.ServeHTTP(rr, req)

		if rr.Code != test.StatusCode {
			t.Errorf(`expected status %d but received %d`,
				test.StatusCode, rr.Code)
		}

		var resp rest.Resp
		json.Unmarshal(rr.Body.Bytes(), &resp)

		if resp.Language != test.ExpectedLanguage {
			t.Errorf(`expected language "%s" but received %s`,
				test.ExpectedLanguage, resp.Language)
		}

		if resp.Translation != test.ExpectedTranslation {
			t.Errorf(`expected Translation "%s" but received %s`,
				test.ExpectedTranslation, resp.Translation)
		}
	}
}
