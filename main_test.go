package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, "expected status code 200")
	expectedBody := "Мир кофе,Сладкоежка"
	require.Equal(t, expectedBody, responseRecorder.Body.String(), "ответ должен содержать первые 2 кафе")
}

func TestMainHandler_WrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=1&city=murmansk", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code, "expected status code 400")
	require.Equal(t, "wrong city", responseRecorder.Body.String(), "expected 'wrong city'")
}
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, "expected status code 200")
	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	require.Equal(t, totalCount, len(list), "количество кафе должно быть равно 4")
}
