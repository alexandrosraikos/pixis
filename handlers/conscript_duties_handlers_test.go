package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/alexandrosraikos/pixis/database"
	"github.com/alexandrosraikos/pixis/models"
	"github.com/gin-gonic/gin"
)

var (
	MockConscriptDuty = models.ConscriptDuty{
		StartTime: time.Now().Add(-1 * time.Hour),
		EndTime:   time.Now().Add(1 * time.Hour),
	}
)

func setupConscriptDutyRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	database.RecreateDatabase("conscript_duties_test.db")
	r := gin.Default()
	r.POST("/conscript_duties", CreateConscriptDuty)
	r.GET("/conscript_duties", GetConscriptDuties)
	r.PUT("/conscript_duties", UpdateConscriptDuty)
	r.DELETE("/conscript_duties", DeleteConscriptDuty)
	return r
}

func beforeEachConscriptDuty(t *testing.T) (*gin.Engine, uint, uint) {
	r := setupConscriptDutyRouter()
	// Create a conscript and a duty for foreign keys
	db := database.GetDB()
	conscript := models.Conscript{
		FirstName:      "DutyTest",
		LastName:       "User",
		RegistryNumber: fmt.Sprintf("cd%d", time.Now().UnixNano()),
		Username:       fmt.Sprintf("cduser%d", time.Now().UnixNano()),
		Password:       "cdpass",
	}
	duty := models.Duty{Label: fmt.Sprintf("DutyForCD%d", time.Now().UnixNano())}
	db.Create(&conscript)
	db.Create(&duty)
	t.Cleanup(func() {})
	return r, conscript.ID, duty.ID
}

func TestCreateConscriptDuty(t *testing.T) {
	r, conscriptID, dutyID := beforeEachConscriptDuty(t)
	cd := MockConscriptDuty
	cd.ConscriptID = conscriptID
	cd.DutyID = dutyID
	jsonValue, _ := json.Marshal(cd)
	req, _ := http.NewRequest("POST", "/conscript_duties", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, w.Code)
	}
}

func TestGetConscriptDutiesByConscript(t *testing.T) {
	r, conscriptID, dutyID := beforeEachConscriptDuty(t)
	cd := MockConscriptDuty
	cd.ConscriptID = conscriptID
	cd.DutyID = dutyID
	jsonValue, _ := json.Marshal(cd)
	req, _ := http.NewRequest("POST", "/conscript_duties", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Use strconv.FormatUint for uint to string conversion
	req, _ = http.NewRequest("GET", "/conscript_duties?conscript_id="+strconv.FormatUint(uint64(conscriptID), 10), nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}
	var cds []models.ConscriptDuty
	err := json.Unmarshal(w.Body.Bytes(), &cds)
	if err != nil {
		t.Errorf("failed to unmarshal response: %v", err)
	}
	if len(cds) == 0 {
		t.Errorf("expected at least one conscript duty, got 0")
	}
}

func TestGetConscriptDutiesByDuty(t *testing.T) {
	r, conscriptID, dutyID := beforeEachConscriptDuty(t)
	cd := MockConscriptDuty
	cd.ConscriptID = conscriptID
	cd.DutyID = dutyID
	jsonValue, _ := json.Marshal(cd)
	req, _ := http.NewRequest("POST", "/conscript_duties", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	req, _ = http.NewRequest("GET", "/conscript_duties?duty_id="+strconv.FormatUint(uint64(dutyID), 10), nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}
	var cds []models.ConscriptDuty
	err := json.Unmarshal(w.Body.Bytes(), &cds)
	if err != nil {
		t.Errorf("failed to unmarshal response: %v", err)
	}
	if len(cds) == 0 {
		t.Errorf("expected at least one conscript duty, got 0")
	}
}

func TestUpdateConscriptDuty(t *testing.T) {
	r, conscriptID, dutyID := beforeEachConscriptDuty(t)
	cd := MockConscriptDuty
	cd.ConscriptID = conscriptID
	cd.DutyID = dutyID
	jsonValue, _ := json.Marshal(cd)
	req, _ := http.NewRequest("POST", "/conscript_duties", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Update EndTime
	update := models.ConscriptDuty{
		ConscriptID: conscriptID,
		DutyID:      dutyID,
		EndTime:     time.Now().Add(2 * time.Hour),
	}
	jsonValue, _ = json.Marshal(update)
	req, _ = http.NewRequest("PUT", "/conscript_duties", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}
	var updated models.ConscriptDuty
	err := json.Unmarshal(w.Body.Bytes(), &updated)
	if err != nil {
		t.Errorf("failed to unmarshal response: %v", err)
	}
	if !updated.EndTime.After(cd.EndTime) {
		t.Errorf("expected EndTime to be updated")
	}
}

func TestDeleteConscriptDuty(t *testing.T) {
	r, conscriptID, dutyID := beforeEachConscriptDuty(t)
	cd := MockConscriptDuty
	cd.ConscriptID = conscriptID
	cd.DutyID = dutyID
	jsonValue, _ := json.Marshal(cd)
	req, _ := http.NewRequest("POST", "/conscript_duties", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Now delete
	deleteBody := struct {
		ConscriptID uint `json:"conscript_id"`
		DutyID      uint `json:"duty_id"`
	}{ConscriptID: conscriptID, DutyID: dutyID}
	jsonValue, _ = json.Marshal(deleteBody)
	req, _ = http.NewRequest("DELETE", "/conscript_duties", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusNoContent {
		t.Errorf("expected status %d, got %d", http.StatusNoContent, w.Code)
	}
}
