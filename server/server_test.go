
package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Approach 1: Direct function testing with mocks
func TestGetMessage(t *testing.T) {
	// Create a mock request
	req := httptest.NewRequest("GET", "/message", nil)
	
	// Create a mock ResponseWriter (ResponseRecorder)
	rr := httptest.NewRecorder()
	
	// Call the function directly
	getMessage(rr, req)
	
	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	
	// Check the response body
	expected := `{"message": "hello"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// Test different HTTP methods
func TestGetMessageDifferentMethods(t *testing.T) {
	tests := []struct {
		name   string
		method string
	}{
		{"GET request", "GET"},
		{"POST request", "POST"},
		{"PUT request", "PUT"},
		{"DELETE request", "DELETE"},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/message", nil)
			rr := httptest.NewRecorder()
			
			getMessage(rr, req)
			
			// Should always return 200 and the same JSON
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
			}
			
			expected := `{"message": "hello"}`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
			}
		})
	}
}

// Approach 2: HTTP Server testing (integration test)
func TestGetMessageHTTPServer(t *testing.T) {
	// Create a test server with your handler
	server := httptest.NewServer(http.HandlerFunc(getMessage))
	defer server.Close() // Important: always close the test server
	
	// Make actual HTTP request to the test server
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()
	
	// Check status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
	
	// Read and check response body
	body := make([]byte, 1024)
	n, _ := resp.Body.Read(body)
	bodyStr := string(body[:n])
	
	expected := `{"message": "hello"}`
	if bodyStr != expected {
		t.Errorf("Expected body %q, got %q", expected, bodyStr)
	}
}

// Test the entire mux (testing multiple endpoints)
func TestServerMux(t *testing.T) {
	// Create the same mux as in main()
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)
	mux.HandleFunc("/message", getMessage)
	
	// Create test server with the mux
	server := httptest.NewServer(mux)
	defer server.Close()
	
	tests := []struct {
		name         string
		path         string
		expectedBody string
	}{
		{"Root endpoint", "/", "This is my website!\n"},
		{"Hello endpoint", "/hello", "Hello!\n"},
		{"Hello with name", "/hello?name=Alice", "Hi Alice! Welcome!\n"},
		{"Message endpoint", "/message", `{"message": "hello"}`},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.Get(server.URL + tt.path)
			if err != nil {
				t.Fatalf("Failed to make request: %v", err)
			}
			defer resp.Body.Close()
			
			body := make([]byte, 1024)
			n, _ := resp.Body.Read(body)
			bodyStr := strings.TrimSpace(string(body[:n]))
			expectedStr := strings.TrimSpace(tt.expectedBody)
			
			if bodyStr != expectedStr {
				t.Errorf("Expected body %q, got %q", expectedStr, bodyStr)
			}
		})
	}
}