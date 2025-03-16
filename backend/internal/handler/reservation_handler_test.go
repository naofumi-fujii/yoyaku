package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/naofumi-fujii/489-yoyaku/backend/internal/model"
	"github.com/naofumi-fujii/489-yoyaku/backend/internal/service"
)

// テスト用なのでインターフェースは削除（handler.goに定義済み）

// モックサービスの実装
type mockReservationService struct {
	createReservationFunc  func(params service.CreateReservationParams) (*model.Reservation, error)
	getAllReservationsFunc func() ([]*model.Reservation, error)
	deleteReservationFunc  func(id string) error
}

func (m *mockReservationService) CreateReservation(params service.CreateReservationParams) (*model.Reservation, error) {
	return m.createReservationFunc(params)
}

func (m *mockReservationService) GetAllReservations() ([]*model.Reservation, error) {
	return m.getAllReservationsFunc()
}

func (m *mockReservationService) DeleteReservation(id string) error {
	return m.deleteReservationFunc(id)
}

func TestCreateReservation(t *testing.T) {
	// Echoのインスタンスを作成
	e := echo.New()

	// モックサービスの準備
	mockSvc := &mockReservationService{
		createReservationFunc: func(params service.CreateReservationParams) (*model.Reservation, error) {
			return model.NewReservation(params.StartTime, params.EndTime), nil
		},
	}

	// ハンドラーの作成
	h := NewReservationHandler(mockSvc)

	// リクエストボディの作成
	now := time.Now()
	startTime := now.Add(time.Hour).Format(time.RFC3339)
	endTime := now.Add(2 * time.Hour).Format(time.RFC3339)
	requestBody := `{"startTime": "` + startTime + `", "endTime": "` + endTime + `"}`

	// リクエストの準備
	req := httptest.NewRequest(http.MethodPost, "/api/reservations", strings.NewReader(requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// ハンドラーを実行
	if err := h.CreateReservation(c); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// レスポンスの検証
	if rec.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, rec.Code)
	}

	// レスポンスボディの検証
	var response model.Reservation
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if response.ID == "" {
		t.Error("Expected ID to be set, got empty string")
	}
}

func TestGetAllReservations(t *testing.T) {
	// Echoのインスタンスを作成
	e := echo.New()

	// モックデータの準備
	now := time.Now()
	expectedReservations := []*model.Reservation{
		model.NewReservation(now, now.Add(time.Hour)),
		model.NewReservation(now.Add(2*time.Hour), now.Add(3*time.Hour)),
	}

	// モックサービスの準備
	mockSvc := &mockReservationService{
		getAllReservationsFunc: func() ([]*model.Reservation, error) {
			return expectedReservations, nil
		},
	}

	// ハンドラーの作成
	h := NewReservationHandler(mockSvc)

	// リクエストの準備
	req := httptest.NewRequest(http.MethodGet, "/api/reservations", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// ハンドラーを実行
	if err := h.GetAllReservations(c); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// レスポンスの検証
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}

	// レスポンスボディの検証
	var response []*model.Reservation
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if len(response) != len(expectedReservations) {
		t.Errorf("Expected %d reservations, got %d", len(expectedReservations), len(response))
	}
}

func TestDeleteReservation(t *testing.T) {
	// Echoのインスタンスを作成
	e := echo.New()

	// テスト用ID
	testID := "test-id"

	// モックサービスの準備
	mockSvc := &mockReservationService{
		deleteReservationFunc: func(id string) error {
			if id != testID {
				t.Errorf("Expected ID %s, got %s", testID, id)
			}
			return nil
		},
	}

	// ハンドラーの作成
	h := NewReservationHandler(mockSvc)

	// リクエストの準備
	req := httptest.NewRequest(http.MethodDelete, "/api/reservations/"+testID, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(testID)

	// ハンドラーを実行
	if err := h.DeleteReservation(c); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// レスポンスの検証
	if rec.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, got %d", http.StatusNoContent, rec.Code)
	}
}