package handler

import (
	"encoding/json"
	"errors"
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

func TestCreateReservation_InvalidRequest(t *testing.T) {
	// Echoのインスタンスを作成
	e := echo.New()

	// ハンドラーの作成
	h := NewReservationHandler(nil) // サービスは使用しないのでnilでOK

	// 不正なリクエストボディの作成
	invalidRequestBody := `{"invalid": "json"`

	// リクエストの準備
	req := httptest.NewRequest(http.MethodPost, "/api/reservations", strings.NewReader(invalidRequestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// ハンドラーを実行
	if err := h.CreateReservation(c); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// レスポンスの検証
	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestCreateReservation_ValidationError(t *testing.T) {
	// Echoのインスタンスを作成
	e := echo.New()

	// ハンドラーの作成
	h := NewReservationHandler(nil) // サービスは使用しないのでnilでOK

	// 必須フィールドが欠けているリクエストボディの作成
	invalidRequestBody := `{"startTime": ""}`

	// リクエストの準備
	req := httptest.NewRequest(http.MethodPost, "/api/reservations", strings.NewReader(invalidRequestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// ハンドラーを実行
	if err := h.CreateReservation(c); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// レスポンスの検証
	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestCreateReservation_InvalidTimeFormat(t *testing.T) {
	// Echoのインスタンスを作成
	e := echo.New()

	// ハンドラーの作成
	h := NewReservationHandler(nil) // サービスは使用しないのでnilでOK

	// 不正な時間フォーマットのリクエストボディ
	invalidTimeFormat := `{"startTime": "invalid-time", "endTime": "2023-01-01T12:00:00Z"}`

	// リクエストの準備
	req := httptest.NewRequest(http.MethodPost, "/api/reservations", strings.NewReader(invalidTimeFormat))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// ハンドラーを実行
	if err := h.CreateReservation(c); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// レスポンスの検証
	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rec.Code)
	}

	// エンドタイムが不正な場合のテスト
	invalidEndTime := `{"startTime": "2023-01-01T12:00:00Z", "endTime": "invalid-time"}`
	req = httptest.NewRequest(http.MethodPost, "/api/reservations", strings.NewReader(invalidEndTime))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	// ハンドラーを実行
	if err := h.CreateReservation(c); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// レスポンスの検証
	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestCreateReservation_EndTimeBeforeStartTime(t *testing.T) {
	// Echoのインスタンスを作成
	e := echo.New()

	// ハンドラーの作成
	h := NewReservationHandler(nil) // サービスは使用しないのでnilでOK

	// 終了時間が開始時間より前のリクエストボディ
	now := time.Now()
	endTime := now.Format(time.RFC3339)
	startTime := now.Add(time.Hour).Format(time.RFC3339)
	invalidRequestBody := `{"startTime": "` + startTime + `", "endTime": "` + endTime + `"}`

	// リクエストの準備
	req := httptest.NewRequest(http.MethodPost, "/api/reservations", strings.NewReader(invalidRequestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// ハンドラーを実行
	if err := h.CreateReservation(c); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// レスポンスの検証
	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestCreateReservation_ServiceError(t *testing.T) {
	// Echoのインスタンスを作成
	e := echo.New()

	// モックサービスの準備 - エラーを返す
	mockSvc := &mockReservationService{
		createReservationFunc: func(params service.CreateReservationParams) (*model.Reservation, error) {
			return nil, errors.New("service error")
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
	if rec.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, rec.Code)
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

func TestGetAllReservations_ServiceError(t *testing.T) {
	// Echoのインスタンスを作成
	e := echo.New()

	// モックサービスの準備 - エラーを返す
	mockSvc := &mockReservationService{
		getAllReservationsFunc: func() ([]*model.Reservation, error) {
			return nil, errors.New("service error")
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
	if rec.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, rec.Code)
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

func TestDeleteReservation_EmptyID(t *testing.T) {
	// Echoのインスタンスを作成
	e := echo.New()

	// ハンドラーの作成
	h := NewReservationHandler(nil) // サービスは使用しないのでnilでOK

	// リクエストの準備 - 空のID
	req := httptest.NewRequest(http.MethodDelete, "/api/reservations/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("") // 空のID

	// ハンドラーを実行
	if err := h.DeleteReservation(c); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// レスポンスの検証
	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestDeleteReservation_ServiceError(t *testing.T) {
	// Echoのインスタンスを作成
	e := echo.New()

	// テスト用ID
	testID := "test-id"

	// モックサービスの準備 - エラーを返す
	mockSvc := &mockReservationService{
		deleteReservationFunc: func(id string) error {
			return errors.New("service error")
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
	if rec.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, rec.Code)
	}
}

func TestRegisterRoutes(t *testing.T) {
	// Echoのインスタンスを作成
	e := echo.New()

	// モックサービスの準備
	mockSvc := &mockReservationService{}

	// ハンドラーの作成
	h := NewReservationHandler(mockSvc)

	// ルート登録
	h.RegisterRoutes(e)

	// 登録されたルートを検証
	routes := e.Routes()
	
	// 期待するルート
	expectedRoutes := []struct {
		path   string
		method string
	}{
		{"/api/reservations", "POST"},
		{"/api/reservations", "GET"},
		{"/api/reservations/:id", "DELETE"},
	}
	
	// ルートが正しく登録されているか確認
	routeFound := make(map[string]bool)
	for _, er := range expectedRoutes {
		for _, route := range routes {
			if route.Path == er.path && route.Method == er.method {
				routeFound[er.path+":"+er.method] = true
				break
			}
		}
	}
	
	// 全てのルートが見つかったか確認
	for _, er := range expectedRoutes {
		key := er.path + ":" + er.method
		if !routeFound[key] {
			t.Errorf("Expected route %s %s to be registered, but it was not found", er.method, er.path)
		}
	}
}