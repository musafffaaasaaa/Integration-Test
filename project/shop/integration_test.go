package shop

import (
	"net/http"
	"net/http/httptest"
	"project/handler"
	"testing"
)

func TestGetProduct(t *testing.T) {

	req, err := http.NewRequest("GET", "/products/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handler.GetProductHandler)

	handler.ServeHTTP(rr, req)

}
