package application_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/sklerakuku/calc-go-lms/internal/application"
)

/*

прописать тесты для handler-ов

прописать конкретные ошибки в файле errors.go (package calculation, var(...))

	ErrInvalidExprassion = errors.New("invalide expration")

прописать файл readme.md


你好， 你的妈！***
*/

func TestCalcHandlerInvalidExpression(t *testing.T) {
	req, err := http.NewRequest("POST", "/api/v1/calculate", bytes.NewBuffer([]byte(`{"expression": "2+a"}`)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	application.CalcHandler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d but got %d", http.StatusBadRequest, w.Code)
	}

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	if response["error"] != "Expression is not valid" {
		t.Errorf("expected error message 'Expression is not valid' but got %v", response["error"])
	}
}

func TestCalcHandlerNonPOSTRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/calculate", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	application.CalcHandler(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("expected status code %d but got %d", http.StatusNotFound, w.Code)
	}

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	if response["error"] != "Internal server error" {
		t.Errorf("expected error message 'Internal server error' but got %v", response["error"])
	}
}

func TestCalcHandlerInvalidJSON(t *testing.T) {
	req, err := http.NewRequest("POST", "/api/v1/calculate", bytes.NewBuffer([]byte(`invalid json`)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	application.CalcHandler(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected status code %d but got %d", http.StatusInternalServerError, w.Code)
	}

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	if response["error"] != "Internal server error" {
		t.Errorf("expected error message 'Internal server error' but got %v", response["error"])
	}
}
func TestHasLetters(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"123", false},
		{"abc", true},
		{"123abc", true},
		{"", false},
	}

	for _, tt := range tests {
		if got := application.HasLetters(tt.input); got != tt.want {
			t.Errorf("hasLetters(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}
func TestConfigFromEnv(t *testing.T) {
	os.Setenv("PORT", "8080")
	config := application.ConfigFromEnv()
	if config.Addr != "8080" {
		t.Errorf("expected port 8080 but got %s", config.Addr)
	}
	os.Unsetenv("PORT")
}

func TestApplicationRunServer(t *testing.T) {
	app := application.New()
	go app.RunServer()

	time.Sleep(100 * time.Millisecond)

	resp, err := http.Get("http://localhost:8080/hello")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "hello\n" {
		t.Errorf("expected body 'hello\\n' but got %s", string(body))
	}
}
