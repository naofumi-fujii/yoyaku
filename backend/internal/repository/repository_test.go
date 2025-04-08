package repository

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/naofumi-fujii/489-yoyaku/backend/internal/model"
)

// Helper function to safely close the database connection
func closeDB(t *testing.T, db *sql.DB) {
	if err := db.Close(); err != nil {
		t.Fatalf("Failed to close database: %v", err)
	}
}

func TestInMemoryReservationRepository_Create(t *testing.T) {
	// 準備
	repo := NewInMemoryReservationRepository()
	now := time.Now()
	reservation := model.NewReservation(now, now.Add(1*time.Hour))

	// 実行
	err := repo.Create(reservation)

	// 検証
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// 保存されたか確認
	stored, err := repo.FindByID(reservation.ID)
	if err != nil {
		t.Errorf("Expected no error when finding reservation, got %v", err)
	}
	if stored == nil {
		t.Error("Expected to find reservation, got nil")
	}
}

func TestInMemoryReservationRepository_FindAll(t *testing.T) {
	// 準備
	repo := NewInMemoryReservationRepository()
	now := time.Now()
	reservation1 := model.NewReservation(now, now.Add(1*time.Hour))
	reservation2 := model.NewReservation(now.Add(2*time.Hour), now.Add(3*time.Hour))

	// 実行：予約を追加
	err := repo.Create(reservation1)
	if err != nil {
		t.Fatalf("Failed to create reservation1: %v", err)
	}
	err = repo.Create(reservation2)
	if err != nil {
		t.Fatalf("Failed to create reservation2: %v", err)
	}

	// 全件取得
	reservations, err := repo.FindAll()

	// 検証
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(reservations) != 2 {
		t.Errorf("Expected 2 reservations, got %d", len(reservations))
	}
}

func TestInMemoryReservationRepository_FindByID(t *testing.T) {
	// 準備
	repo := NewInMemoryReservationRepository()
	now := time.Now()
	reservation := model.NewReservation(now, now.Add(1*time.Hour))
	err := repo.Create(reservation)
	if err != nil {
		t.Fatalf("Failed to create reservation: %v", err)
	}

	// テストケース
	tests := []struct {
		name    string
		id      string
		wantNil bool
	}{
		{
			name:    "存在するID",
			id:      reservation.ID,
			wantNil: false,
		},
		{
			name:    "存在しないID",
			id:      "non-existent-id",
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 実行
			got, err := repo.FindByID(tt.id)

			// 検証
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			if (got == nil) != tt.wantNil {
				t.Errorf("FindByID() got = %v, want nil = %v", got, tt.wantNil)
			}
			if !tt.wantNil && got.ID != tt.id {
				t.Errorf("Expected ID %s, got %s", tt.id, got.ID)
			}
		})
	}
}

func TestInMemoryReservationRepository_Delete(t *testing.T) {
	// 準備
	repo := NewInMemoryReservationRepository()
	now := time.Now()
	reservation := model.NewReservation(now, now.Add(1*time.Hour))
	err := repo.Create(reservation)
	if err != nil {
		t.Fatalf("Failed to create reservation: %v", err)
	}

	// 実行：存在するIDを削除
	err = repo.Delete(reservation.ID)

	// 検証
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// 削除されたか確認
	found, _ := repo.FindByID(reservation.ID)
	if found != nil {
		t.Error("Expected reservation to be deleted, but it still exists")
	}

	// 実行：存在しないIDを削除（エラーにならないことを確認）
	err = repo.Delete("non-existent-id")

	// 検証
	if err != nil {
		t.Errorf("Expected no error when deleting non-existent ID, got %v", err)
	}
}

// MySQLレポジトリのテスト

func TestMySQLReservationRepository_New(t *testing.T) {
	// SQLMockのセットアップ
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer closeDB(t, db)

	// テーブル作成クエリの期待値を設定
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS reservations").WillReturnResult(sqlmock.NewResult(0, 0))

	// レポジトリの作成
	repo, err := NewMySQLReservationRepository(db)

	// 検証
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if repo == nil {
		t.Error("Expected repository to be created, got nil")
	}

	// モックの期待通りに実行されたか確認
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestMySQLReservationRepository_New_Error(t *testing.T) {
	// SQLMockのセットアップ
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer closeDB(t, db)

	// テーブル作成クエリでエラーを返すように設定
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS reservations").WillReturnError(errors.New("database error"))

	// レポジトリの作成
	repo, err := NewMySQLReservationRepository(db)

	// 検証
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if repo != nil {
		t.Errorf("Expected nil repository, got %v", repo)
	}

	// モックの期待通りに実行されたか確認
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestMySQLReservationRepository_Create(t *testing.T) {
	// SQLMockのセットアップ
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer closeDB(t, db)

	// テーブル作成クエリの期待値を設定（NewMySQLReservationRepositoryでの呼び出し）
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS reservations").WillReturnResult(sqlmock.NewResult(0, 0))

	// レポジトリの作成
	repo, err := NewMySQLReservationRepository(db)
	if err != nil {
		t.Fatalf("Failed to create repository: %v", err)
	}

	// テスト対象の予約データ
	now := time.Now()
	reservation := model.NewReservation(now, now.Add(1*time.Hour))

	// INSERTクエリの期待値を設定
	mock.ExpectExec("INSERT INTO reservations").WithArgs(
		reservation.ID,
		reservation.StartTime,
		reservation.EndTime,
		reservation.CreatedAt,
		reservation.UpdatedAt,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	// 実行
	err = repo.Create(reservation)

	// 検証
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// モックの期待通りに実行されたか確認
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestMySQLReservationRepository_Create_Error(t *testing.T) {
	// SQLMockのセットアップ
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer closeDB(t, db)

	// テーブル作成クエリの期待値を設定（NewMySQLReservationRepositoryでの呼び出し）
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS reservations").WillReturnResult(sqlmock.NewResult(0, 0))

	// レポジトリの作成
	repo, err := NewMySQLReservationRepository(db)
	if err != nil {
		t.Fatalf("Failed to create repository: %v", err)
	}

	// テスト対象の予約データ
	now := time.Now()
	reservation := model.NewReservation(now, now.Add(1*time.Hour))

	// INSERTクエリでエラーを返すように設定
	mock.ExpectExec("INSERT INTO reservations").WithArgs(
		reservation.ID,
		reservation.StartTime,
		reservation.EndTime,
		reservation.CreatedAt,
		reservation.UpdatedAt,
	).WillReturnError(errors.New("database error"))

	// 実行
	err = repo.Create(reservation)

	// 検証
	if err == nil {
		t.Error("Expected error, got nil")
	}

	// モックの期待通りに実行されたか確認
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestMySQLReservationRepository_FindAll(t *testing.T) {
	// SQLMockのセットアップ
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer closeDB(t, db)

	// テーブル作成クエリの期待値を設定（NewMySQLReservationRepositoryでの呼び出し）
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS reservations").WillReturnResult(sqlmock.NewResult(0, 0))

	// レポジトリの作成
	repo, err := NewMySQLReservationRepository(db)
	if err != nil {
		t.Fatalf("Failed to create repository: %v", err)
	}

	// テストデータ
	now := time.Now()
	id1 := uuid.New().String()
	id2 := uuid.New().String()
	startTime1 := now
	endTime1 := now.Add(1 * time.Hour)
	startTime2 := now.Add(2 * time.Hour)
	endTime2 := now.Add(3 * time.Hour)
	createdAt := now
	updatedAt := now

	// SELECTクエリの結果を設定
	rows := sqlmock.NewRows([]string{"id", "start_time", "end_time", "created_at", "updated_at"}).
		AddRow(id1, startTime1, endTime1, createdAt, updatedAt).
		AddRow(id2, startTime2, endTime2, createdAt, updatedAt)

	// SELECTクエリの期待値を設定
	mock.ExpectQuery("SELECT id, start_time, end_time, created_at, updated_at FROM reservations").
		WillReturnRows(rows)

	// 実行
	reservations, err := repo.FindAll()

	// 検証
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(reservations) != 2 {
		t.Errorf("Expected 2 reservations, got %d", len(reservations))
	}
	if reservations[0].ID != id1 {
		t.Errorf("Expected first reservation ID %s, got %s", id1, reservations[0].ID)
	}
	if reservations[1].ID != id2 {
		t.Errorf("Expected second reservation ID %s, got %s", id2, reservations[1].ID)
	}

	// モックの期待通りに実行されたか確認
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestMySQLReservationRepository_FindAll_QueryError(t *testing.T) {
	// SQLMockのセットアップ
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer closeDB(t, db)

	// テーブル作成クエリの期待値を設定（NewMySQLReservationRepositoryでの呼び出し）
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS reservations").WillReturnResult(sqlmock.NewResult(0, 0))

	// レポジトリの作成
	repo, err := NewMySQLReservationRepository(db)
	if err != nil {
		t.Fatalf("Failed to create repository: %v", err)
	}

	// SELECTクエリでエラーを返すように設定
	mock.ExpectQuery("SELECT id, start_time, end_time, created_at, updated_at FROM reservations").
		WillReturnError(errors.New("database error"))

	// 実行
	reservations, err := repo.FindAll()

	// 検証
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if reservations != nil {
		t.Errorf("Expected nil reservations, got %v", reservations)
	}

	// モックの期待通りに実行されたか確認
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestMySQLReservationRepository_FindAll_ScanError(t *testing.T) {
	// SQLMockのセットアップ
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer closeDB(t, db)

	// テーブル作成クエリの期待値を設定（NewMySQLReservationRepositoryでの呼び出し）
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS reservations").WillReturnResult(sqlmock.NewResult(0, 0))

	// レポジトリの作成
	repo, err := NewMySQLReservationRepository(db)
	if err != nil {
		t.Fatalf("Failed to create repository: %v", err)
	}

	// 型不一致によるスキャンエラーを発生させるために不正な列タイプを設定
	rows := sqlmock.NewRows([]string{"id", "start_time", "end_time", "created_at", "updated_at"}).
		AddRow("id1", "not-a-time", "not-a-time", "not-a-time", "not-a-time")

	// SELECTクエリの期待値を設定
	mock.ExpectQuery("SELECT id, start_time, end_time, created_at, updated_at FROM reservations").
		WillReturnRows(rows)

	// 実行
	reservations, err := repo.FindAll()

	// 検証
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if reservations != nil {
		t.Errorf("Expected nil reservations, got %v", reservations)
	}

	// モックの期待通りに実行されたか確認
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestMySQLReservationRepository_FindByID(t *testing.T) {
	// SQLMockのセットアップ
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer closeDB(t, db)

	// テーブル作成クエリの期待値を設定（NewMySQLReservationRepositoryでの呼び出し）
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS reservations").WillReturnResult(sqlmock.NewResult(0, 0))

	// レポジトリの作成
	repo, err := NewMySQLReservationRepository(db)
	if err != nil {
		t.Fatalf("Failed to create repository: %v", err)
	}

	// テストデータ
	now := time.Now()
	id := uuid.New().String()
	startTime := now
	endTime := now.Add(1 * time.Hour)
	createdAt := now
	updatedAt := now

	// SELECTクエリの結果を設定
	rows := sqlmock.NewRows([]string{"id", "start_time", "end_time", "created_at", "updated_at"}).
		AddRow(id, startTime, endTime, createdAt, updatedAt)

	// SELECTクエリの期待値を設定
	mock.ExpectQuery("SELECT id, start_time, end_time, created_at, updated_at FROM reservations WHERE id = ?").
		WithArgs(id).
		WillReturnRows(rows)

	// 実行
	reservation, err := repo.FindByID(id)

	// 検証
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if reservation == nil {
		t.Error("Expected reservation, got nil")
		return
	}
	if reservation.ID != id {
		t.Errorf("Expected reservation ID %s, got %s", id, reservation.ID)
	}

	// モックの期待通りに実行されたか確認
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestMySQLReservationRepository_FindByID_NotFound(t *testing.T) {
	// SQLMockのセットアップ
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer closeDB(t, db)

	// テーブル作成クエリの期待値を設定（NewMySQLReservationRepositoryでの呼び出し）
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS reservations").WillReturnResult(sqlmock.NewResult(0, 0))

	// レポジトリの作成
	repo, err := NewMySQLReservationRepository(db)
	if err != nil {
		t.Fatalf("Failed to create repository: %v", err)
	}

	// テストデータ
	id := uuid.New().String()

	// SELECTクエリで行が見つからないことを設定
	mock.ExpectQuery("SELECT id, start_time, end_time, created_at, updated_at FROM reservations WHERE id = ?").
		WithArgs(id).
		WillReturnError(sql.ErrNoRows)

	// 実行
	reservation, err := repo.FindByID(id)

	// 検証
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if reservation != nil {
		t.Errorf("Expected nil reservation, got %v", reservation)
	}

	// モックの期待通りに実行されたか確認
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestMySQLReservationRepository_FindByID_Error(t *testing.T) {
	// SQLMockのセットアップ
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer closeDB(t, db)

	// テーブル作成クエリの期待値を設定（NewMySQLReservationRepositoryでの呼び出し）
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS reservations").WillReturnResult(sqlmock.NewResult(0, 0))

	// レポジトリの作成
	repo, err := NewMySQLReservationRepository(db)
	if err != nil {
		t.Fatalf("Failed to create repository: %v", err)
	}

	// テストデータ
	id := uuid.New().String()

	// SELECTクエリでエラーを返すように設定
	mock.ExpectQuery("SELECT id, start_time, end_time, created_at, updated_at FROM reservations WHERE id = ?").
		WithArgs(id).
		WillReturnError(errors.New("database error"))

	// 実行
	reservation, err := repo.FindByID(id)

	// 検証
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if reservation != nil {
		t.Errorf("Expected nil reservation, got %v", reservation)
	}

	// モックの期待通りに実行されたか確認
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestMySQLReservationRepository_Delete(t *testing.T) {
	// SQLMockのセットアップ
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer closeDB(t, db)

	// テーブル作成クエリの期待値を設定（NewMySQLReservationRepositoryでの呼び出し）
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS reservations").WillReturnResult(sqlmock.NewResult(0, 0))

	// レポジトリの作成
	repo, err := NewMySQLReservationRepository(db)
	if err != nil {
		t.Fatalf("Failed to create repository: %v", err)
	}

	// テストデータ
	id := uuid.New().String()

	// DELETEクエリの期待値を設定
	mock.ExpectExec("DELETE FROM reservations WHERE id = ?").
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(0, 1))

	// 実行
	err = repo.Delete(id)

	// 検証
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// モックの期待通りに実行されたか確認
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestMySQLReservationRepository_Delete_Error(t *testing.T) {
	// SQLMockのセットアップ
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer closeDB(t, db)

	// テーブル作成クエリの期待値を設定（NewMySQLReservationRepositoryでの呼び出し）
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS reservations").WillReturnResult(sqlmock.NewResult(0, 0))

	// レポジトリの作成
	repo, err := NewMySQLReservationRepository(db)
	if err != nil {
		t.Fatalf("Failed to create repository: %v", err)
	}

	// テストデータ
	id := uuid.New().String()

	// DELETEクエリでエラーを返すように設定
	mock.ExpectExec("DELETE FROM reservations WHERE id = ?").
		WithArgs(id).
		WillReturnError(errors.New("database error"))

	// 実行
	err = repo.Delete(id)

	// 検証
	if err == nil {
		t.Error("Expected error, got nil")
	}

	// モックの期待通りに実行されたか確認
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}