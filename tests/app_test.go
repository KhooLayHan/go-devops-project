package app_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KhooLayHan/go-devops-project/internal/app"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.Handler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "Welcome to DevOps Project 1 with Go!\n", rr.Body.String())
}