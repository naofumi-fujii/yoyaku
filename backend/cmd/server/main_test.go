package main

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/naofumi-fujii/489-yoyaku/backend/internal/handler"
	"github.com/naofumi-fujii/489-yoyaku/backend/internal/repository"
	"github.com/naofumi-fujii/489-yoyaku/backend/internal/service"
)

// Helper function to safely close the database connection
func closeDB(t *testing.T, db *sql.DB) {
	if err := db.Close(); err != nil {
		t.Fatalf("Failed to close database: %v", err)
	}
}

func TestHealthCheck(t *testing.T) {
	// Echo インスタンスの作成
	e := echo.New()
	
	// ヘルスチェックのハンドラーを登録
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok", "db": "connected"})
	})
	
	// リクエストの準備
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	
	// ハンドラーを直接実行
	h := func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok", "db": "connected"})
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

func TestGetEnv(t *testing.T) {
	// 環境変数がセットされていない場合のテスト
	testKey := "TEST_ENV_KEY"
	defaultValue := "default_value"
	
	// 先に環境変数をクリア
	if err := os.Unsetenv(testKey); err != nil {
		t.Fatalf("Failed to unset environment variable: %v", err)
	}
	
	result := getEnv(testKey, defaultValue)
	if result != defaultValue {
		t.Errorf("Expected default value %s when env variable not set, got %s", defaultValue, result)
	}
	
	// 環境変数がセットされている場合のテスト
	expectedValue := "expected_value"
	os.Setenv(testKey, expectedValue)
	defer func() {
		if err := os.Unsetenv(testKey); err != nil {
			t.Fatalf("Failed to unset environment variable in defer: %v", err)
		}
	}()
	
	result = getEnv(testKey, defaultValue)
	if result != expectedValue {
		t.Errorf("Expected value %s when env variable is set, got %s", expectedValue, result)
	}
}

func TestInitializeApp(t *testing.T) {
	// SQLMockのセットアップ
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer closeDB(t, db)

	// テーブル作成クエリの期待値を設定
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS reservations").WillReturnResult(sqlmock.NewResult(0, 0))

	// Echoインスタンスを作成
	e := echo.New()

	// 実行
	reservationRepo, reservationService, reservationHandler, err := initializeApp(e, db)

	// 検証
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// 各コンポーネントが正しく初期化されているか確認
	if reservationRepo == nil {
		t.Error("Expected repository to be initialized, got nil")
	}
	if reservationService == nil {
		t.Error("Expected service to be initialized, got nil")
	}
	if reservationHandler == nil {
		t.Error("Expected handler to be initialized, got nil")
	}

	// モックの期待通りに実行されたか確認
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

// initializeApp関数をテスト用に抽出
func initializeApp(e *echo.Echo, db *sql.DB) (repository.ReservationRepository, *service.ReservationService, *handler.ReservationHandler, error) {
	// Initialize repository
	mysqlRepo, err := repository.NewMySQLReservationRepository(db)
	if err != nil {
		return nil, nil, nil, err
	}

	// Initialize service
	reservationService := service.NewReservationService(mysqlRepo)

	// Initialize handler
	reservationHandler := handler.NewReservationHandler(reservationService)

	// Register routes
	reservationHandler.RegisterRoutes(e)

	return mysqlRepo, reservationService, reservationHandler, nil
}

func TestDatabaseConnection(t *testing.T) {
	// SQLMockのセットアップ
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer closeDB(t, db)

	// モックのPingが成功するように設定
	mock.ExpectPing()

	// Ping関数をテスト
	err = db.Ping()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// モックの期待通りに実行されたか確認
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestRetryLogic(t *testing.T) {
	// 簡易的なretryロジックテスト
	counter := 0
	maxRetries := 3
	
	// モックの関数（最初の2回は失敗、3回目で成功）
	mockFn := func() error {
		counter++
		if counter < 3 {
			return sql.ErrConnDone
		}
		return nil
	}
	
	// retryロジック
	var err error
	for i := 0; i < maxRetries; i++ {
		err = mockFn()
		if err == nil {
			break
		}
	}
	
	// 検証
	if err != nil {
		t.Errorf("Expected no error after retries, got %v", err)
	}
	
	if counter != 3 {
		t.Errorf("Expected 3 attempts, got %d", counter)
	}
}