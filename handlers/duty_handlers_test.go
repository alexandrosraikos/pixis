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

var MockDuty = models.Duty{
	Label: "TestDuty",
}

func setupDutyRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	database.RecreateDatabase("duty_test.db")
	r := gin.Default()
	r.POST("/duties", CreateDuty)
	r.GET("/duties", GetDuties)
	r.GET("/duties/:id", GetDuty)
	r.PUT("/duties/:id", UpdateDuty)
	r.DELETE("/duties/:id", DeleteDuty)
	return r
}

func beforeEachDuty(t *testing.T) *gin.Engine {
	r := setupDutyRouter()
	t.Cleanup(func() {})
	return r
}

func TestCreateDuty(t *testing.T) {
	r := beforeEachDuty(t)
	duty := MockDuty
	jsonValue, _ := json.Marshal(duty)
	req, _ := http.NewRequest("POST", "/duties", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, w.Code)
	}
}

func TestGetDuties(t *testing.T) {
	r := beforeEachDuty(t)
	duty := MockDuty
	duty.Label = "AnotherDuty"
	jsonValue, _ := json.Marshal(duty)
	req, _ := http.NewRequest("POST", "/duties", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	req, _ = http.NewRequest("GET", "/duties", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}
	var duties []models.Duty
	err := json.Unmarshal(w.Body.Bytes(), &duties)
	if err != nil {
		t.Errorf("failed to unmarshal response: %v", err)
	}
	if len(duties) == 0 {
		t.Errorf("expected at least one duty, got 0")
	}
}

func TestGetDuty(t *testing.T) {
	r := beforeEachDuty(t)
	duty := MockDuty
	duty.Label = "UniqueDuty"
	jsonValue, _ := json.Marshal(duty)
	req, _ := http.NewRequest("POST", "/duties", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var created models.Duty
	json.Unmarshal(w.Body.Bytes(), &created)

	url := fmt.Sprintf("/duties/%d", created.ID)
	req, _ = http.NewRequest("GET", url, nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}
	var got models.Duty
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Errorf("failed to unmarshal response: %v", err)
	}
	if got.ID != created.ID {
		t.Errorf("expected ID %d, got %d", created.ID, got.ID)
	}
}

func TestUpdateDuty(t *testing.T) {
	r := beforeEachDuty(t)
	duty := MockDuty
	duty.Label = "DutyToUpdate"
	jsonValue, _ := json.Marshal(duty)
	req, _ := http.NewRequest("POST", "/duties", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var created models.Duty
	json.Unmarshal(w.Body.Bytes(), &created)

	update := models.Duty{Label: "UpdatedDuty"}
	jsonValue, _ = json.Marshal(update)
	url := fmt.Sprintf("/duties/%d", created.ID)
	req, _ = http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}
	var updated models.Duty
	json.Unmarshal(w.Body.Bytes(), &updated)
	if updated.Label != "UpdatedDuty" {
		t.Errorf("expected Label 'UpdatedDuty', got '%s'", updated.Label)
	}
}

func TestDeleteDuty(t *testing.T) {
	r := beforeEachDuty(t)
	duty := MockDuty
	duty.Label = "DutyToDelete"
	jsonValue, _ := json.Marshal(duty)
	req, _ := http.NewRequest("POST", "/duties", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var created models.Duty
	json.Unmarshal(w.Body.Bytes(), &created)

	url := fmt.Sprintf("/duties/%d", created.ID)
	req, _ = http.NewRequest("DELETE", url, nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}
}
