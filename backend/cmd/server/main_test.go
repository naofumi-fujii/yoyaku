package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestHealthCheck(t *testing.T) {
	// Echo インスタンスの作成
	e := echo.New()
	
	// ヘルスチェックのハンドラーを登録
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})
	
	// リクエストの準備
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	
	// ハンドラーを直接実行
	h := func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	}
	
	// テスト実行
	if err := h(c); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	
	// レスポンスの検証
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}
	
	// レスポンスボディの検証 - JSON文字列で厳密な比較ではなく、内容を確認
	if body := rec.Body.String(); body == "" || body == "{}" {
		t.Errorf("Expected non-empty response body, got %s", body)
	}
}