package main

import (
    "net/http"
    "net/http/httptest"
    "os"
    "testing"
)

// TestGreetHandler_ValidRequest tests the /greet endpoint with a valid GET request
func TestGreetHandler_ValidRequest(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/greet", nil)
    w := httptest.NewRecorder()
    
    GreetHandler(w, req)
    
    res := w.Result()
    if res.StatusCode != http.StatusOK {
        t.Errorf("expected status OK; got %v", res.Status)
    }
    
    expected := `{"message":"Hello, welcome to the Go Greet API!"}`
    body := w.Body.String()
    if body != expected {
        t.Errorf("expected body %v; got %v", expected, body)
    }
}

// TestGreetHandler_InvalidMethod tests the /greet endpoint with an invalid HTTP method (POST)
func TestGreetHandler_InvalidMethod(t *testing.T) {
    req := httptest.NewRequest(http.MethodPost, "/greet", nil)
    w := httptest.NewRecorder()
    
    GreetHandler(w, req)
    
    res := w.Result()
    if res.StatusCode != http.StatusMethodNotAllowed {
        t.Errorf("expected status Method Not Allowed; got %v", res.Status)
    }

    expected := "Method Not Allowed\n"
    body := w.Body.String()
    if body != expected {
        t.Errorf("expected body %v; got %v", expected, body)
    }
}

// TestGetPort_DefaultPort tests that the default port 8080 is used if the PORT environment variable is not set
func TestGetPort_DefaultPort(t *testing.T) {
    os.Unsetenv("PORT") // Ensure PORT is not set
    port := GetPort()
    if port != "8080" {
        t.Errorf("expected default port 8080; got %v", port)
    }
}

// TestGetPort_CustomPort tests that a custom port is returned when PORT environment variable is set
func TestGetPort_CustomPort(t *testing.T) {
    os.Setenv("PORT", "9090")
    defer os.Unsetenv("PORT") // Clean up after test
    port := GetPort()
    if port != "9090" {
        t.Errorf("expected custom port 9090; got %v", port)
    }
}
