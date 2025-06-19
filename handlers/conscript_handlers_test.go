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

var MockConscript = models.Conscript{
	FirstName:      "TestFirst",
	LastName:       "TestLast",
	RegistryNumber: "99999",
	Username:       "testuser",
	Password:       "testpass",
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	database.RecreateDatabase("conscript_test.db")
	r := gin.Default()
	r.POST("/conscripts", CreateConscript)
	r.GET("/conscripts", GetConscripts)
	r.GET("/conscripts/:id", GetConscript)
	r.PUT("/conscripts/:id", UpdateConscript)
	r.DELETE("/conscripts/:id", DeleteConscript)
	return r
}

func beforeEach(t *testing.T) (*gin.Engine, uint) {
	r := setupRouter()
	// Create a department for foreign key
	db := database.GetDB()
	dept := MockDepartment

	// Ensure department with same label does not exist
	db.Where("label = ?", dept.Label).Delete(&models.Department{})

	if err := db.Create(&dept).Error; err != nil {
		t.Fatalf("failed to create department: %v", err)
	}
	t.Cleanup(func() {
		// Add DB cleanup code here if needed
	})
	return r, dept.ID
}

func TestCreateConscript(t *testing.T) {
	r, deptID := beforeEach(t)
	conscript := MockConscript
	conscript.DepartmentID = deptID
	jsonValue, _ := json.Marshal(conscript)
	req, _ := http.NewRequest("POST", "/conscripts", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, w.Code)
	}
	var created models.Conscript
	json.Unmarshal(w.Body.Bytes(), &created)
	if created.DepartmentID != deptID {
		t.Errorf("expected DepartmentID %d, got %d", deptID, created.DepartmentID)
	}
	if created.CreatedAt.IsZero() || created.UpdatedAt.IsZero() {
		t.Errorf("expected CreatedAt and UpdatedAt to be set")
	}
}

func TestGetConscripts(t *testing.T) {
	r, deptID := beforeEach(t)
	conscript := MockConscript
	conscript.RegistryNumber = "54321"
	conscript.Username = "janesmith"
	conscript.DepartmentID = deptID
	jsonValue, _ := json.Marshal(conscript)
	req, _ := http.NewRequest("POST", "/conscripts", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

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
	for _, c := range conscripts {
		if c.DepartmentID != deptID {
			t.Errorf("expected DepartmentID %d, got %d", deptID, c.DepartmentID)
		}
	}
}

func TestGetConscript(t *testing.T) {
	r, deptID := beforeEach(t)
	conscript := MockConscript
	conscript.RegistryNumber = "11111"
	conscript.Username = "alice"
	conscript.DepartmentID = deptID
	jsonValue, _ := json.Marshal(conscript)
	req, _ := http.NewRequest("POST", "/conscripts", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var created models.Conscript
	json.Unmarshal(w.Body.Bytes(), &created)

	url := fmt.Sprintf("/conscripts/%d", created.ID)
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
	if got.DepartmentID != deptID {
		t.Errorf("expected DepartmentID %d, got %d", deptID, got.DepartmentID)
	}
}

func TestUpdateConscript(t *testing.T) {
	r, deptID := beforeEach(t)
	conscript := MockConscript
	conscript.RegistryNumber = "22222"
	conscript.Username = "bob"
	conscript.DepartmentID = deptID
	jsonValue, _ := json.Marshal(conscript)
	req, _ := http.NewRequest("POST", "/conscripts", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var created models.Conscript
	json.Unmarshal(w.Body.Bytes(), &created)

	update := models.Conscript{FirstName: "Robert", DepartmentID: deptID}
	jsonValue, _ = json.Marshal(update)
	url := fmt.Sprintf("/conscripts/%d", created.ID)
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
	if updated.DepartmentID != deptID {
		t.Errorf("expected DepartmentID %d, got %d", deptID, updated.DepartmentID)
	}
}

func TestDeleteConscript(t *testing.T) {
	r, deptID := beforeEach(t)
	conscript := MockConscript
	conscript.RegistryNumber = "33333"
	conscript.Username = "carl"
	conscript.DepartmentID = deptID
	jsonValue, _ := json.Marshal(conscript)
	req, _ := http.NewRequest("POST", "/conscripts", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var created models.Conscript
	json.Unmarshal(w.Body.Bytes(), &created)

	url := fmt.Sprintf("/conscripts/%d", created.ID)
	req, _ = http.NewRequest("DELETE", url, nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusNoContent {
		t.Errorf("expected status %d, got %d", http.StatusNoContent, w.Code)
	}
}
