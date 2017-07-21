package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"strconv"
)

func TestHealth(t *testing.T) {
	router := GetMappedRouter()
	req, _ := http.NewRequest("GET", "/health", nil)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, req)

	if writer.Code != http.StatusOK {
		println("Status NOK: " + strconv.Itoa(writer.Code))
		t.Fail()
	}

	if writer.Body.String() != "fuck yeah!" {
		println("Body NOK: " + writer.Body.String())
		t.Fail()
	}
}
