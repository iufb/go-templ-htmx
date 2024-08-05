package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/iufb/go-templ-htmx/types"
)

func TestUserServiceHandler(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)
	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			Email:    "invalid",
			Password: "leejieun",
		}
		RouteTestRunner(t, payload, http.StatusBadRequest, handler)
	})
	t.Run("should  correctly register user", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			Email:    "valid@gmail.com",
			Password: "leejieun",
		}
		RouteTestRunner(t, payload, http.StatusCreated, handler)
	})
}

func RouteTestRunner(t *testing.T, payload types.RegisterUserPayload, expectedCode int, handler *Handler) {
	marshalled, _ := json.Marshal(payload)
	req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/register", handler.handleRegister)
	router.ServeHTTP(rr, req)
	if rr.Code != expectedCode {
		t.Errorf("Expected status code %d, got %d", expectedCode, rr.Code)
	}
}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) GetUserById(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(user types.User) error {
	return nil
}
