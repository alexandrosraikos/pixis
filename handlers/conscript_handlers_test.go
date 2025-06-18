package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alexandrosraikos/pixis/database"
	"github.com/alexandrosraikos/pixis/models"
	"github.com/gin-gonic/gin"
)

// MockConscript is a reusable mock for tests
var MockConscript = models.Conscript{
	FirstName:      "TestFirst",
	LastName:       "TestLast",
	RegistryNumber: "99999",
	Username:       "testuser",
	Password:       "testpass",
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	database.ConnectDatabase("conscript_test.db")
	r := gin.Default()
	r.POST("/conscripts", CreateConscript)
	r.GET("/conscripts", GetConscripts)
	r.GET("/conscripts/:id", GetConscript)
	r.PUT("/conscripts/:id", UpdateConscript)
	r.DELETE("/conscripts/:id", DeleteConscript)
	return r
}

func beforeEach(t *testing.T) *gin.Engine {
	// Optionally, clean up the DB file here if you want a fresh DB for each test
	r := setupRouter()
	t.Cleanup(func() {
		// Add DB cleanup code here if needed
	})
	return r
}

func TestCreateConscript(t *testing.T) {
	r := beforeEach(t)
	conscript := MockConscript
	jsonValue, _ := json.Marshal(conscript)
	req, _ := http.NewRequest("POST", "/conscripts", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, w.Code)
	}
}

func TestGetConscripts(t *testing.T) {
	r := beforeEach(t)
	conscript := MockConscript
	conscript.RegistryNumber = "54321"
	conscript.Username = "janesmith"
	jsonValue, _ := json.Marshal(conscript)
	req, _ := http.NewRequest("POST", "/conscripts", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Now test GET /conscripts
	req, _ = http.NewRequest("GET", "/conscripts", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}
	var conscripts []models.Conscript
	err := json.Unmarshal(w.Body.Bytes(), &conscripts)
	if err != nil {
		t.Errorf("failed to unmarshal response: %v", err)
	}
	if len(conscripts) == 0 {
		t.Errorf("expected at least one conscript, got 0")
	}
}

func TestGetConscript(t *testing.T) {
	r := beforeEach(t)
	conscript := MockConscript
	conscript.RegistryNumber = "11111"
	conscript.Username = "alice"
	jsonValue, _ := json.Marshal(conscript)
	req, _ := http.NewRequest("POST", "/conscripts", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var created models.Conscript
	json.Unmarshal(w.Body.Bytes(), &created)

	// Now test GET /conscripts/:id
	url := "/conscripts/" + json.Number(fmt.Sprintf("%d", created.ID)).String()
	req, _ = http.NewRequest("GET", url, nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}
	var got models.Conscript
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Errorf("failed to unmarshal response: %v", err)
	}
	if got.ID != created.ID {
		t.Errorf("expected ID %d, got %d", created.ID, got.ID)
	}
}

func TestUpdateConscript(t *testing.T) {
	r := beforeEach(t)
	conscript := MockConscript
	conscript.RegistryNumber = "22222"
	conscript.Username = "bob"
	jsonValue, _ := json.Marshal(conscript)
	req, _ := http.NewRequest("POST", "/conscripts", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var created models.Conscript
	json.Unmarshal(w.Body.Bytes(), &created)

	// Now test PUT /conscripts/:id
	update := models.Conscript{FirstName: "Robert"}
	jsonValue, _ = json.Marshal(update)
	url := "/conscripts/" + json.Number(fmt.Sprintf("%d", created.ID)).String()
	req, _ = http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}
	var updated models.Conscript
	json.Unmarshal(w.Body.Bytes(), &updated)
	if updated.FirstName != "Robert" {
		t.Errorf("expected FirstName 'Robert', got '%s'", updated.FirstName)
	}
}

func TestDeleteConscript(t *testing.T) {
	r := beforeEach(t)
	conscript := MockConscript
	conscript.RegistryNumber = "33333"
	conscript.Username = "carl"
	jsonValue, _ := json.Marshal(conscript)
	req, _ := http.NewRequest("POST", "/conscripts", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var created models.Conscript
	json.Unmarshal(w.Body.Bytes(), &created)

	// Now test DELETE /conscripts/:id
	url := "/conscripts/" + json.Number(fmt.Sprintf("%d", created.ID)).String()
	req, _ = http.NewRequest("DELETE", url, nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusNoContent {
		t.Errorf("expected status %d, got %d", http.StatusNoContent, w.Code)
	}
}
