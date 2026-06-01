package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockAuthService para evitar depender da implementação real do serviço
type mockAuthService struct{}

func (m *mockAuthService) Register(ctx context.Context, email, password string) error {
	if email == "existente@test.com" {
		return errors.New("usuário já existe")
	}
	return nil
}

func (m *mockAuthService) Login(ctx context.Context, email, password string) (string, error) {
	if email == "valido@test.com" && password == "123" {
		return "mocked-jwt-token", nil
	}
	return "", errors.New("invalid")
}

func TestAuthHandler_Login(t *testing.T) {
	mockSvc := &mockAuthService{}
	handler := NewAuthHandler(mockSvc)

	reqBody, _ := json.Marshal(authRequest{Email: "valido@test.com", Password: "123"})
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(reqBody))
	rr := httptest.NewRecorder()

	handler.Login(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("esperava status %d, obteve %d", http.StatusOK, rr.Code)
	}

	var response map[string]string
	json.Unmarshal(rr.Body.Bytes(), &response)

	if response["token"] != "mocked-jwt-token" {
		t.Errorf("esperava token 'mocked-jwt-token', obteve '%s'", response["token"])
	}
}
