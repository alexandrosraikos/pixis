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

// MockDepartment is a reusable mock for tests
var MockDepartment = models.Department{
	Label: "Test Department",
}

func setupDepartmentRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	database.RecreateDatabase("department_test.db")
	r := gin.Default()
	r.POST("/departments", CreateDepartment)
	r.GET("/departments", GetDepartments)
	r.GET("/departments/:id", GetDepartment)
	r.PUT("/departments/:id", UpdateDepartment)
	r.DELETE("/departments/:id", DeleteDepartment)
	return r
}

func beforeEachDepartment(t *testing.T) *gin.Engine {
	r := setupDepartmentRouter()
	t.Cleanup(func() {
		// Add DB cleanup code here if needed
	})
	return r
}

func TestCreateDepartment(t *testing.T) {
	r := beforeEachDepartment(t)
	department := MockDepartment
	jsonValue, _ := json.Marshal(department)
	req, _ := http.NewRequest("POST", "/departments", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, w.Code)
	}
}

func TestGetDepartments(t *testing.T) {
	r := beforeEachDepartment(t)
	department := MockDepartment
	department.Label = "Another Department"
	jsonValue, _ := json.Marshal(department)
	req, _ := http.NewRequest("POST", "/departments", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	req, _ = http.NewRequest("GET", "/departments", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}
	var departments []models.Department
	err := json.Unmarshal(w.Body.Bytes(), &departments)
	if err != nil {
		t.Errorf("failed to unmarshal response: %v", err)
	}
	if len(departments) == 0 {
		t.Errorf("expected at least one department, got 0")
	}
}

func TestGetDepartment(t *testing.T) {
	r := beforeEachDepartment(t)
	department := MockDepartment
	department.Label = "Unique Department"
	jsonValue, _ := json.Marshal(department)
	req, _ := http.NewRequest("POST", "/departments", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var created models.Department
	json.Unmarshal(w.Body.Bytes(), &created)

	url := fmt.Sprintf("/departments/%d", created.ID)
	req, _ = http.NewRequest("GET", url, nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}
	var got models.Department
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Errorf("failed to unmarshal response: %v", err)
	}
	if got.ID != created.ID {
		t.Errorf("expected ID %d, got %d", created.ID, got.ID)
	}
}

func TestUpdateDepartment(t *testing.T) {
	r := beforeEachDepartment(t)
	department := MockDepartment
	department.Label = "Department To Update"
	jsonValue, _ := json.Marshal(department)
	req, _ := http.NewRequest("POST", "/departments", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var created models.Department
	json.Unmarshal(w.Body.Bytes(), &created)

	update := models.Department{Label: "Updated Department"}
	jsonValue, _ = json.Marshal(update)
	url := fmt.Sprintf("/departments/%d", created.ID)
	req, _ = http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}
	var updated models.Department
	json.Unmarshal(w.Body.Bytes(), &updated)
	if updated.Label != "Updated Department" {
		t.Errorf("expected Label 'Updated Department', got '%s'", updated.Label)
	}
}

func TestDeleteDepartment(t *testing.T) {
	r := beforeEachDepartment(t)
	department := MockDepartment
	department.Label = "Department To Delete"
	jsonValue, _ := json.Marshal(department)
	req, _ := http.NewRequest("POST", "/departments", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var created models.Department
	json.Unmarshal(w.Body.Bytes(), &created)

	url := fmt.Sprintf("/departments/%d", created.ID)
	req, _ = http.NewRequest("DELETE", url, nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}
}
