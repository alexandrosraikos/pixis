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

var MockService = models.Service{
	Label: "TestService",
}

func setupServiceRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	database.RecreateDatabase("service_test.db")
	r := gin.Default()
	r.POST("/services", CreateService)
	r.GET("/services", GetServices)
	r.GET("/services/:id", GetService)
	r.PUT("/services/:id", UpdateService)
	r.DELETE("/services/:id", DeleteService)
	return r
}

func beforeEachService(t *testing.T) *gin.Engine {
	r := setupServiceRouter()
	t.Cleanup(func() {})
	return r
}

func TestCreateService(t *testing.T) {
	r := beforeEachService(t)
	service := MockService
	jsonValue, _ := json.Marshal(service)
	req, _ := http.NewRequest("POST", "/services", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, w.Code)
	}
}

func TestGetServices(t *testing.T) {
	r := beforeEachService(t)
	service := MockService
	service.Label = "AnotherService"
	jsonValue, _ := json.Marshal(service)
	req, _ := http.NewRequest("POST", "/services", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	req, _ = http.NewRequest("GET", "/services", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}
	var services []models.Service
	err := json.Unmarshal(w.Body.Bytes(), &services)
	if err != nil {
		t.Errorf("failed to unmarshal response: %v", err)
	}
	if len(services) == 0 {
		t.Errorf("expected at least one service, got 0")
	}
}

func TestGetService(t *testing.T) {
	r := beforeEachService(t)
	service := MockService
	service.Label = "UniqueService"
	jsonValue, _ := json.Marshal(service)
	req, _ := http.NewRequest("POST", "/services", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var created models.Service
	json.Unmarshal(w.Body.Bytes(), &created)

	url := fmt.Sprintf("/services/%d", created.ID)
	req, _ = http.NewRequest("GET", url, nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}
	var got models.Service
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Errorf("failed to unmarshal response: %v", err)
	}
	if got.ID != created.ID {
		t.Errorf("expected ID %d, got %d", created.ID, got.ID)
	}
}

func TestUpdateService(t *testing.T) {
	r := beforeEachService(t)
	service := MockService
	service.Label = "ServiceToUpdate"
	jsonValue, _ := json.Marshal(service)
	req, _ := http.NewRequest("POST", "/services", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var created models.Service
	json.Unmarshal(w.Body.Bytes(), &created)

	update := models.Service{Label: "UpdatedService"}
	jsonValue, _ = json.Marshal(update)
	url := fmt.Sprintf("/services/%d", created.ID)
	req, _ = http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}
	var updated models.Service
	json.Unmarshal(w.Body.Bytes(), &updated)
	if updated.Label != "UpdatedService" {
		t.Errorf("expected Label 'UpdatedService', got '%s'", updated.Label)
	}
}

func TestDeleteService(t *testing.T) {
	r := beforeEachService(t)
	service := MockService
	service.Label = "ServiceToDelete"
	jsonValue, _ := json.Marshal(service)
	req, _ := http.NewRequest("POST", "/services", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var created models.Service
	json.Unmarshal(w.Body.Bytes(), &created)

	url := fmt.Sprintf("/services/%d", created.ID)
	req, _ = http.NewRequest("DELETE", url, nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}
}
