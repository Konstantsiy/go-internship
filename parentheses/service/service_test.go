package service

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProcessRequest(t *testing.T) {
	const URL = "http://localhost:8080/generate?n="

	testTable := []struct {
		Param              string
		ExpectedStatusCode int
	}{
		{
			Param:              "123",
			ExpectedStatusCode: http.StatusOK,
		},
		{
			Param:              "-34",
			ExpectedStatusCode: http.StatusBadRequest,
		},
		{
			Param:              "3dffdfs4",
			ExpectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testTable {
		req, err := http.NewRequest("GET", URL+tc.Param, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(ProcessRequest)
		handler.ServeHTTP(rr, req)

		if rr.Code != tc.ExpectedStatusCode {
			t.Errorf("handler returned wrong status code: want %v got %v", tc.ExpectedStatusCode, rr.Code)
		}
	}
}
