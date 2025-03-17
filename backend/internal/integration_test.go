package internal

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/naofumi-fujii/489-yoyaku/backend/internal/handler"
	"github.com/naofumi-fujii/489-yoyaku/backend/internal/model"
	"github.com/naofumi-fujii/489-yoyaku/backend/internal/repository"
	"github.com/naofumi-fujii/489-yoyaku/backend/internal/service"
)

func setupTest() *echo.Echo {
	e := echo.New()
	
	// リポジトリ、サービス、ハンドラーの初期化
	repo := repository.NewInMemoryReservationRepository()
	svc := service.NewReservationService(repo)
	h := handler.NewReservationHandler(svc)
	
	// ルートの登録
	h.RegisterRoutes(e)
	
	return e
}

func TestIntegrationCreateAndGetReservation(t *testing.T) {
	// テスト用サーバーのセットアップ
	e := setupTest()
	
	// 予約作成リクエストの準備
	now := time.Now()
	startTime := now.Add(time.Hour).Format(time.RFC3339)
	endTime := now.Add(2 * time.Hour).Format(time.RFC3339)
	payload := map[string]string{
		"startTime": startTime,
		"endTime":   endTime,
	}
	
	payloadBytes, _ := json.Marshal(payload)
	
	// 予約作成リクエストの送信
	req := httptest.NewRequest(http.MethodPost, "/api/reservations", bytes.NewReader(payloadBytes))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	
	e.ServeHTTP(rec, req)
	
	// レスポンスの検証
	if rec.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, rec.Code)
	}
	
	// 作成された予約のIDを取得
	var createdReservation model.Reservation
	if err := json.Unmarshal(rec.Body.Bytes(), &createdReservation); err != nil {
		t.Fatalf("Failed to unmarshal created reservation: %v", err)
	}
	
	// 予約一覧取得リクエストの送信
	req = httptest.NewRequest(http.MethodGet, "/api/reservations", nil)
	rec = httptest.NewRecorder()
	
	e.ServeHTTP(rec, req)
	
	// レスポンスの検証
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}
	
	// 予約一覧の検証
	var reservations []model.Reservation
	if err := json.Unmarshal(rec.Body.Bytes(), &reservations); err != nil {
		t.Fatalf("Failed to unmarshal reservations: %v", err)
	}
	
	if len(reservations) != 1 {
		t.Errorf("Expected 1 reservation, got %d", len(reservations))
	}
	
	if reservations[0].ID != createdReservation.ID {
		t.Errorf("Expected reservation ID %s, got %s", createdReservation.ID, reservations[0].ID)
	}
}

func TestIntegrationDeleteReservation(t *testing.T) {
	// テスト用サーバーのセットアップ
	e := setupTest()
	
	// 予約作成リクエストの準備
	now := time.Now()
	startTime := now.Add(time.Hour).Format(time.RFC3339)
	endTime := now.Add(2 * time.Hour).Format(time.RFC3339)
	payload := map[string]string{
		"startTime": startTime,
		"endTime":   endTime,
	}
	
	payloadBytes, _ := json.Marshal(payload)
	
	// 予約作成リクエストの送信
	req := httptest.NewRequest(http.MethodPost, "/api/reservations", bytes.NewReader(payloadBytes))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	
	e.ServeHTTP(rec, req)
	
	// 作成された予約のIDを取得
	var createdReservation model.Reservation
	if err := json.Unmarshal(rec.Body.Bytes(), &createdReservation); err != nil {
		t.Fatalf("Failed to unmarshal created reservation: %v", err)
	}
	
	// 予約削除リクエストの送信
	req = httptest.NewRequest(http.MethodDelete, "/api/reservations/"+createdReservation.ID, nil)
	rec = httptest.NewRecorder()
	
	e.ServeHTTP(rec, req)
	
	// レスポンスの検証
	if rec.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, got %d", http.StatusNoContent, rec.Code)
	}
	
	// 予約が削除されたことを確認
	req = httptest.NewRequest(http.MethodGet, "/api/reservations", nil)
	rec = httptest.NewRecorder()
	
	e.ServeHTTP(rec, req)
	
	// 予約一覧の検証
	var reservations []model.Reservation
	if err := json.Unmarshal(rec.Body.Bytes(), &reservations); err != nil {
		t.Fatalf("Failed to unmarshal reservations: %v", err)
	}
	
	if len(reservations) != 0 {
		t.Errorf("Expected 0 reservations after deletion, got %d", len(reservations))
	}
}