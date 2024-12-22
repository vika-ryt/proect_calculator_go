package calculation_test

import (
	"io"
	"net/http"
	"testing"

	"net/http/httptest"
)

func TestCalcHandlerSuccessCase(t *testing.T) {
	expected := make(map[string]string)
	expected["result"] = "5"
	req := httptest.NewRequest(http.MethodGet, " http://google.com", nil)
	w := httptest.NewRecorder()
	CalcHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if string(data[1]) != expected["result"] {
		t.Errorf("Something was wrong")
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("wrong status code")
	}
}

func CalcHandler(w *httptest.ResponseRecorder, req *http.Request) {
	panic("unimplemented")
}

func TestRequestHandlerBadRequestCase(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/greet", nil)
	w := httptest.NewRecorder()
	CalcHandler(w, req)
	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("wrong status code")
	}
}